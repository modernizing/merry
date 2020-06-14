d3.json("/manifest-map.json").then(function(data){
    console.log(data);
    var width = 900;
    var height = 1104;

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

    var margin = ({top: 20, right: 20, bottom: 20, left: 100});
    var y = d3.scalePoint(nodes.map(d => d.id).sort(d3.ascending), [margin.top, height - margin.bottom])
    var color = d3.scaleOrdinal(nodes.map(d => d.group).sort(d3.ascending), d3.schemeCategory10)
    var step = 14;
    height = (nodes.length - 1) * step + margin.top + margin.bottom
    var svg = d3.select("#svg").append("svg")
        .attr("width", width)
        .attr("height", height);

    const label = svg.append("g")
        .attr("font-family", "sans-serif")
        .attr("font-size", 10)
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
        })
        .on("mouseout", d => {
            svg.classed("hover", false);
            label.classed("primary", false);
            label.classed("secondary", false);
            path.classed("primary", false).order();
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
});
