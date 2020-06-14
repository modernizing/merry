function renderVertical() {
    return function (data) {
        var width = 1440;
        var height = 2048;

        const nodes = data.nodes.map(({id, group}) => ({
            id,
            sourceLinks: [],
            targetLinks: [],
            group
        }));

        const nodeById = new Map(nodes.map(d => [d.id, d]));

        const links = data.links.map(({source, target, value}) => ({
            source: nodeById.get(source),
            target: nodeById.get(target),
            value
        }));

        for (const link of links) {
            const {source, target, value} = link;
            if (target && source) {
                source.sourceLinks.push(link);
                target.targetLinks.push(link);
            }
        }

        var margin = ({top: 20, right: 20, bottom: 20, left: 250});
        var step = 14;

        height = (nodes.length - 1) * step + margin.top + margin.bottom
        var y = d3.scalePoint(nodes.map(d => d.id).sort(d3.ascending), [margin.top, height - margin.bottom])
        var color = d3.scaleOrdinal(nodes.map(d => d.group).sort(d3.ascending), d3.schemeCategory10)
        var svg = d3.select("#svg").append("svg")
            .attr("width", width)
            .attr("height", height);

        // Define the div for the tooltip
        var div = d3.select("body").append("div")
            .attr("class", "tooltip")
            .style("opacity", 0);

        const label = svg.append("g")
            .attr("font-family", "sans-serif")
            .attr("font-size", 12)
            .attr("text-anchor", "end")
            .selectAll("g")
            .data(nodes)
            .join("g")
            .attr("transform", d => `translate(${margin.left},${d.y = y(d.id)})`)
            .call(g => g.append("text")
                .attr("x", -6)
                .attr("dy", "0.35em")
                .attr("fill", d => d3.lab(color(d.group)).darker(2))
                .text(d => d.id))
            .call(g => g.append("circle")
                .attr("r", 3)
                .attr("fill", d => color(d.group)));

        const path = svg.insert("g", "*")
            .attr("fill", "none")
            .attr("stroke-opacity", 0.6)
            .attr("stroke-width", 1.5)
            .selectAll("path")
            .data(links)
            .join("path")
            .attr("stroke", d => d.source.group === d.target.group ? color(d.source.group) : "#aaa")
            .attr("d", arc);

        const overlay = svg.append("g")
            .attr("fill", "none")
            .attr("pointer-events", "all")
            .selectAll("rect")
            .data(nodes)
            .join("rect")
            .attr("width", margin.left + 40)
            .attr("height", step)
            .attr("y", d => y(d.id) - step / 2)
            .on("mouseover", d => {
                svg.classed("hover", true);
                label.classed("primary", n => n === d);
                label.classed("secondary", n => n.sourceLinks.some(l => l.target === d) || n.targetLinks.some(l => l.source === d));
                path.classed("primary", l => l.source === d || l.target === d).filter(".primary").raise();

                var result = "<h3>" + d.id + "</h3>";
                for (let linkElement of d.targetLinks) {
                    result += "usedby: " + linkElement.source.id + "<br>"
                }

                for (let linkElement of d.sourceLinks) {
                    result += "use: " + linkElement.target.id + "<br>"
                }

                div.transition()
                    .duration(200)
                    .style("opacity", .9);
                div.html(result)
                    .style("display", "block")
                    .style("left", (d3.event.pageX) + "px")
                    .style("top", (d3.event.pageY - 28) + "px");
            })
            .on("mouseout", d => {
                svg.classed("hover", false);
                label.classed("primary", false);
                label.classed("secondary", false);
                path.classed("primary", false).order();
                div.style("display", "none")
            });

        function update(order) {
            y.domain(nodes.sort(order.value).map(d => d.id));

            const t = svg.transition()
                .duration(750);

            label.transition(t)
                .delay((d, i) => i * 20)
                .attrTween("transform", d => {
                    const i = d3.interpolateNumber(d.y, y(d.id));
                    return t => `translate(${margin.left},${d.y = i(t)})`;
                });

            path.transition(t)
                .duration(750 + nodes.length * 20)
                .attrTween("d", d => () => arc(d));

            overlay.transition(t)
                .delay((d, i) => i * 20)
                .attr("y", d => y(d.id) - step / 2);
        }

        function arc(d) {
            const y1 = d.source.y;
            const y2 = d.target.y;
            const r = Math.abs(y2 - y1) / 2;
            return `M${margin.left},${y1}A${r},${r} 0,0,${y1 < y2 ? 1 : 0} ${margin.left},${y2}`;
        }

        d3.select("#selectSort").on("change", function () {
            var value = d3.select("#selectSort").node().value;
            var selectMap = {
                Group: {
                    value: (a, b) => a.group - b.group || d3.ascending(a.id, b.id)
                },
                Name: {
                    value: (a, b) => d3.ascending(a.id, b.id)
                },
                Frequency: {
                    value: (a, b) => d3.sum(b.sourceLinks, l => l.value) + d3.sum(b.targetLinks, l => l.value) - d3.sum(a.sourceLinks, l => l.value) - d3.sum(a.targetLinks, l => l.value) || d3.ascending(a.id, b.id)
                }
            }

            let select = selectMap[value];
            update(select);
        })
    };
}

d3.json("/manifest-map.json").then(renderVertical());
// d3.json("output.json").then(renderVertical());

function renderCircle() {
    return function () {
        function hierarchy(data, delimiter = ".") {
            let root;
            const map = new Map;
            data.forEach(function find(data) {
                const {name} = data;
                if (map.has(name)) return map.get(name);
                const i = name.lastIndexOf(delimiter);
                map.set(name, data);
                if (i >= 0) {
                    find({name: name.substring(0, i), children: []}).children.push(data);
                    data.name = name.substring(i + 1);
                } else {
                    root = data;
                }
                return data;
            });
            return root;
        }
        var jdata = [
            {"name":"flare.vis.events.DataEvent","size":2313,"imports":["flare.vis.events.SelectionEvent"]},
            {"name":"flare.vis.events.SelectionEvent","size":1880,"imports":["flare.vis.events.DataEvent"]}
            ]

        var data = hierarchy(jdata)
        function bilink(root) {
            const map = new Map(root.leaves().map(d => [id(d), d]));
            for (const d of root.leaves()) {
                d.incoming = []
                d.outgoing = d.data.imports.map(i => [d, map.get(i)]);
            }
            for (const d of root.leaves()) {
                for (const o of d.outgoing) {
                    o[1].incoming.push(o)
                }
            }
            return root;
        }

        function id(node) {
            return `${node.parent ? id(node.parent) + "." : ""}${node.data.name}`;
        }

        var colorin = "#00f",
            colorout = "#f00",
            colornone = "#ccc",
            width = 954,
            radius = width / 2,
            line = d3.lineRadial()
                .curve(d3.curveBundle.beta(0.85))
                .radius(d => d.y)
                .angle(d => d.x),
            tree = d3.cluster()
                .size([2 * Math.PI, radius - 100]);

        const root = tree(bilink(d3.hierarchy(data)
            .sort((a, b) => d3.ascending(a.height, b.height) || d3.ascending(a.data.name, b.data.name))));

        var svg = d3.select("#circle").append("svg")
            .attr("viewBox", [-width / 2, -width / 2, width, width]);

        const node = svg.append("g")
            .attr("font-family", "sans-serif")
            .attr("font-size", 10)
            .selectAll("g")
            .data(root.leaves())
            .join("g")
            .attr("transform", d => `rotate(${d.x * 180 / Math.PI - 90}) translate(${d.y},0)`)
            .append("text")
            .attr("dy", "0.31em")
            .attr("x", d => d.x < Math.PI ? 6 : -6)
            .attr("text-anchor", d => d.x < Math.PI ? "start" : "end")
            .attr("transform", d => d.x >= Math.PI ? "rotate(180)" : null)
            .text(d => d.data.name)
            .each(function (d) {
                d.text = this;
            })
            .on("mouseover", overed)
            .on("mouseout", outed)
            .call(text => text.append("title").text(d => `${id(d)}
${d.outgoing.length} outgoing
${d.incoming.length} incoming`));

        const link = svg.append("g")
            .attr("stroke", colornone)
            .attr("fill", "none")
            .selectAll("path")
            .data(root.leaves().flatMap(leaf => leaf.outgoing))
            .join("path")
            .style("mix-blend-mode", "multiply")
            .attr("d", ([i, o]) => line(i.path(o)))
            .each(function (d) {
                d.path = this;
            });

        function overed(d) {
            link.style("mix-blend-mode", null);
            d3.select(this).attr("font-weight", "bold");
            d3.selectAll(d.incoming.map(d => d.path)).attr("stroke", colorin).raise();
            d3.selectAll(d.incoming.map(([d]) => d.text)).attr("fill", colorin).attr("font-weight", "bold");
            d3.selectAll(d.outgoing.map(d => d.path)).attr("stroke", colorout).raise();
            d3.selectAll(d.outgoing.map(([, d]) => d.text)).attr("fill", colorout).attr("font-weight", "bold");
        }

        function outed(d) {
            link.style("mix-blend-mode", "multiply");
            d3.select(this).attr("font-weight", null);
            d3.selectAll(d.incoming.map(d => d.path)).attr("stroke", null);
            d3.selectAll(d.incoming.map(([d]) => d.text)).attr("fill", null).attr("font-weight", null);
            d3.selectAll(d.outgoing.map(d => d.path)).attr("stroke", null);
            d3.selectAll(d.outgoing.map(([, d]) => d.text)).attr("fill", null).attr("font-weight", null);
        }

    }
}

// d3.json("output.json").then(renderCircle());

