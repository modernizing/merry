// Code generated from Manifest.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 46, 553,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 359, 10, 3, 12, 3,
	14, 3, 362, 11, 3, 3, 3, 3, 3, 3, 3, 7, 3, 367, 10, 3, 12, 3, 14, 3, 370,
	11, 3, 5, 3, 372, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 377, 10, 4, 12, 4, 14,
	4, 380, 11, 4, 3, 5, 3, 5, 6, 5, 384, 10, 5, 13, 5, 14, 5, 385, 6, 5, 388,
	10, 5, 13, 5, 14, 5, 389, 3, 5, 5, 5, 393, 10, 5, 3, 6, 3, 6, 3, 6, 6,
	6, 398, 10, 6, 13, 6, 14, 6, 399, 3, 7, 3, 7, 7, 7, 404, 10, 7, 12, 7,
	14, 7, 407, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 7, 42, 487, 10, 42, 12,
	42, 14, 42, 490, 11, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 5, 44, 497,
	10, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 7, 45, 506, 10,
	45, 12, 45, 14, 45, 509, 11, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3,
	46, 5, 46, 517, 10, 46, 3, 46, 5, 46, 520, 10, 46, 3, 46, 3, 46, 3, 46,
	6, 46, 525, 10, 46, 13, 46, 14, 46, 526, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 5, 46, 534, 10, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 5, 49,
	542, 10, 49, 3, 50, 3, 50, 7, 50, 546, 10, 50, 12, 50, 14, 50, 549, 11,
	50, 3, 50, 5, 50, 552, 10, 50, 2, 2, 51, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67,
	35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85,
	44, 87, 45, 89, 46, 91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 3, 2, 14, 4, 2,
	42, 43, 47, 48, 3, 2, 67, 92, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34,
	6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 10, 2, 36, 36, 41, 41, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 3, 2, 50, 53, 3, 2, 50, 57,
	5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4,
	2, 50, 59, 97, 97, 2, 585, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 3, 101, 3,
	2, 2, 2, 5, 371, 3, 2, 2, 2, 7, 373, 3, 2, 2, 2, 9, 392, 3, 2, 2, 2, 11,
	394, 3, 2, 2, 2, 13, 401, 3, 2, 2, 2, 15, 408, 3, 2, 2, 2, 17, 410, 3,
	2, 2, 2, 19, 412, 3, 2, 2, 2, 21, 414, 3, 2, 2, 2, 23, 416, 3, 2, 2, 2,
	25, 418, 3, 2, 2, 2, 27, 420, 3, 2, 2, 2, 29, 422, 3, 2, 2, 2, 31, 424,
	3, 2, 2, 2, 33, 426, 3, 2, 2, 2, 35, 428, 3, 2, 2, 2, 37, 430, 3, 2, 2,
	2, 39, 432, 3, 2, 2, 2, 41, 434, 3, 2, 2, 2, 43, 436, 3, 2, 2, 2, 45, 438,
	3, 2, 2, 2, 47, 440, 3, 2, 2, 2, 49, 443, 3, 2, 2, 2, 51, 446, 3, 2, 2,
	2, 53, 449, 3, 2, 2, 2, 55, 452, 3, 2, 2, 2, 57, 455, 3, 2, 2, 2, 59, 458,
	3, 2, 2, 2, 61, 461, 3, 2, 2, 2, 63, 464, 3, 2, 2, 2, 65, 466, 3, 2, 2,
	2, 67, 468, 3, 2, 2, 2, 69, 470, 3, 2, 2, 2, 71, 472, 3, 2, 2, 2, 73, 474,
	3, 2, 2, 2, 75, 476, 3, 2, 2, 2, 77, 478, 3, 2, 2, 2, 79, 480, 3, 2, 2,
	2, 81, 482, 3, 2, 2, 2, 83, 484, 3, 2, 2, 2, 85, 493, 3, 2, 2, 2, 87, 496,
	3, 2, 2, 2, 89, 502, 3, 2, 2, 2, 91, 533, 3, 2, 2, 2, 93, 535, 3, 2, 2,
	2, 95, 537, 3, 2, 2, 2, 97, 541, 3, 2, 2, 2, 99, 543, 3, 2, 2, 2, 101,
	102, 7, 120, 2, 2, 102, 103, 7, 103, 2, 2, 103, 104, 7, 116, 2, 2, 104,
	105, 7, 117, 2, 2, 105, 106, 7, 107, 2, 2, 106, 107, 7, 113, 2, 2, 107,
	108, 7, 112, 2, 2, 108, 4, 3, 2, 2, 2, 109, 110, 7, 79, 2, 2, 110, 111,
	7, 99, 2, 2, 111, 112, 7, 112, 2, 2, 112, 113, 7, 107, 2, 2, 113, 114,
	7, 104, 2, 2, 114, 115, 7, 103, 2, 2, 115, 116, 7, 117, 2, 2, 116, 117,
	7, 118, 2, 2, 117, 118, 7, 47, 2, 2, 118, 119, 7, 88, 2, 2, 119, 120, 7,
	103, 2, 2, 120, 121, 7, 116, 2, 2, 121, 122, 7, 117, 2, 2, 122, 123, 7,
	107, 2, 2, 123, 124, 7, 113, 2, 2, 124, 372, 7, 112, 2, 2, 125, 126, 7,
	68, 2, 2, 126, 127, 7, 119, 2, 2, 127, 128, 7, 112, 2, 2, 128, 129, 7,
	102, 2, 2, 129, 130, 7, 110, 2, 2, 130, 131, 7, 103, 2, 2, 131, 132, 7,
	47, 2, 2, 132, 133, 7, 67, 2, 2, 133, 134, 7, 101, 2, 2, 134, 135, 7, 118,
	2, 2, 135, 136, 7, 107, 2, 2, 136, 137, 7, 120, 2, 2, 137, 138, 7, 99,
	2, 2, 138, 139, 7, 118, 2, 2, 139, 140, 7, 113, 2, 2, 140, 372, 7, 116,
	2, 2, 141, 142, 7, 69, 2, 2, 142, 143, 7, 116, 2, 2, 143, 144, 7, 103,
	2, 2, 144, 145, 7, 99, 2, 2, 145, 146, 7, 118, 2, 2, 146, 147, 7, 103,
	2, 2, 147, 148, 7, 102, 2, 2, 148, 149, 7, 47, 2, 2, 149, 150, 7, 68, 2,
	2, 150, 372, 7, 123, 2, 2, 151, 152, 7, 75, 2, 2, 152, 153, 7, 111, 2,
	2, 153, 154, 7, 114, 2, 2, 154, 155, 7, 113, 2, 2, 155, 156, 7, 116, 2,
	2, 156, 157, 7, 118, 2, 2, 157, 158, 7, 47, 2, 2, 158, 159, 7, 82, 2, 2,
	159, 160, 7, 99, 2, 2, 160, 161, 7, 101, 2, 2, 161, 162, 7, 109, 2, 2,
	162, 163, 7, 99, 2, 2, 163, 164, 7, 105, 2, 2, 164, 372, 7, 103, 2, 2,
	165, 166, 7, 71, 2, 2, 166, 167, 7, 122, 2, 2, 167, 168, 7, 114, 2, 2,
	168, 169, 7, 113, 2, 2, 169, 170, 7, 116, 2, 2, 170, 171, 7, 118, 2, 2,
	171, 172, 7, 47, 2, 2, 172, 173, 7, 82, 2, 2, 173, 174, 7, 99, 2, 2, 174,
	175, 7, 101, 2, 2, 175, 176, 7, 109, 2, 2, 176, 177, 7, 99, 2, 2, 177,
	178, 7, 105, 2, 2, 178, 372, 7, 103, 2, 2, 179, 180, 7, 68, 2, 2, 180,
	181, 7, 119, 2, 2, 181, 182, 7, 112, 2, 2, 182, 183, 7, 102, 2, 2, 183,
	184, 7, 110, 2, 2, 184, 185, 7, 103, 2, 2, 185, 186, 7, 47, 2, 2, 186,
	187, 7, 88, 2, 2, 187, 188, 7, 103, 2, 2, 188, 189, 7, 116, 2, 2, 189,
	190, 7, 117, 2, 2, 190, 191, 7, 107, 2, 2, 191, 192, 7, 113, 2, 2, 192,
	372, 7, 112, 2, 2, 193, 194, 7, 68, 2, 2, 194, 195, 7, 119, 2, 2, 195,
	196, 7, 112, 2, 2, 196, 197, 7, 102, 2, 2, 197, 198, 7, 110, 2, 2, 198,
	199, 7, 103, 2, 2, 199, 200, 7, 47, 2, 2, 200, 201, 7, 80, 2, 2, 201, 202,
	7, 99, 2, 2, 202, 203, 7, 111, 2, 2, 203, 372, 7, 103, 2, 2, 204, 205,
	7, 68, 2, 2, 205, 206, 7, 119, 2, 2, 206, 207, 7, 112, 2, 2, 207, 208,
	7, 102, 2, 2, 208, 209, 7, 110, 2, 2, 209, 210, 7, 103, 2, 2, 210, 211,
	7, 47, 2, 2, 211, 212, 7, 70, 2, 2, 212, 213, 7, 103, 2, 2, 213, 214, 7,
	117, 2, 2, 214, 215, 7, 101, 2, 2, 215, 216, 7, 116, 2, 2, 216, 217, 7,
	107, 2, 2, 217, 218, 7, 114, 2, 2, 218, 219, 7, 118, 2, 2, 219, 220, 7,
	107, 2, 2, 220, 221, 7, 113, 2, 2, 221, 372, 7, 112, 2, 2, 222, 223, 7,
	68, 2, 2, 223, 224, 7, 119, 2, 2, 224, 225, 7, 112, 2, 2, 225, 226, 7,
	102, 2, 2, 226, 227, 7, 110, 2, 2, 227, 228, 7, 103, 2, 2, 228, 229, 7,
	47, 2, 2, 229, 230, 7, 70, 2, 2, 230, 231, 7, 113, 2, 2, 231, 232, 7, 101,
	2, 2, 232, 233, 7, 87, 2, 2, 233, 234, 7, 84, 2, 2, 234, 372, 7, 78, 2,
	2, 235, 236, 7, 68, 2, 2, 236, 237, 7, 119, 2, 2, 237, 238, 7, 112, 2,
	2, 238, 239, 7, 102, 2, 2, 239, 240, 7, 110, 2, 2, 240, 241, 7, 103, 2,
	2, 241, 242, 7, 47, 2, 2, 242, 243, 7, 79, 2, 2, 243, 244, 7, 99, 2, 2,
	244, 245, 7, 112, 2, 2, 245, 246, 7, 107, 2, 2, 246, 247, 7, 104, 2, 2,
	247, 248, 7, 103, 2, 2, 248, 249, 7, 117, 2, 2, 249, 250, 7, 118, 2, 2,
	250, 251, 7, 88, 2, 2, 251, 252, 7, 103, 2, 2, 252, 253, 7, 116, 2, 2,
	253, 254, 7, 117, 2, 2, 254, 255, 7, 107, 2, 2, 255, 256, 7, 113, 2, 2,
	256, 372, 7, 112, 2, 2, 257, 258, 7, 68, 2, 2, 258, 259, 7, 119, 2, 2,
	259, 260, 7, 112, 2, 2, 260, 261, 7, 102, 2, 2, 261, 262, 7, 110, 2, 2,
	262, 263, 7, 103, 2, 2, 263, 264, 7, 47, 2, 2, 264, 265, 7, 88, 2, 2, 265,
	266, 7, 103, 2, 2, 266, 267, 7, 112, 2, 2, 267, 268, 7, 102, 2, 2, 268,
	269, 7, 113, 2, 2, 269, 372, 7, 116, 2, 2, 270, 271, 7, 68, 2, 2, 271,
	272, 7, 119, 2, 2, 272, 273, 7, 112, 2, 2, 273, 274, 7, 102, 2, 2, 274,
	275, 7, 110, 2, 2, 275, 276, 7, 103, 2, 2, 276, 277, 7, 47, 2, 2, 277,
	278, 7, 85, 2, 2, 278, 279, 7, 123, 2, 2, 279, 280, 7, 111, 2, 2, 280,
	281, 7, 100, 2, 2, 281, 282, 7, 113, 2, 2, 282, 283, 7, 110, 2, 2, 283,
	284, 7, 107, 2, 2, 284, 285, 7, 101, 2, 2, 285, 286, 7, 80, 2, 2, 286,
	287, 7, 99, 2, 2, 287, 288, 7, 111, 2, 2, 288, 372, 7, 103, 2, 2, 289,
	290, 7, 75, 2, 2, 290, 291, 7, 111, 2, 2, 291, 292, 7, 114, 2, 2, 292,
	293, 7, 110, 2, 2, 293, 294, 7, 103, 2, 2, 294, 295, 7, 111, 2, 2, 295,
	296, 7, 103, 2, 2, 296, 297, 7, 112, 2, 2, 297, 298, 7, 118, 2, 2, 298,
	299, 7, 99, 2, 2, 299, 300, 7, 118, 2, 2, 300, 301, 7, 107, 2, 2, 301,
	302, 7, 113, 2, 2, 302, 303, 7, 112, 2, 2, 303, 304, 7, 47, 2, 2, 304,
	305, 7, 88, 2, 2, 305, 306, 7, 103, 2, 2, 306, 307, 7, 116, 2, 2, 307,
	308, 7, 117, 2, 2, 308, 309, 7, 107, 2, 2, 309, 310, 7, 113, 2, 2, 310,
	372, 7, 112, 2, 2, 311, 312, 7, 75, 2, 2, 312, 313, 7, 111, 2, 2, 313,
	314, 7, 114, 2, 2, 314, 315, 7, 110, 2, 2, 315, 316, 7, 103, 2, 2, 316,
	317, 7, 111, 2, 2, 317, 318, 7, 103, 2, 2, 318, 319, 7, 112, 2, 2, 319,
	320, 7, 118, 2, 2, 320, 321, 7, 99, 2, 2, 321, 322, 7, 118, 2, 2, 322,
	323, 7, 107, 2, 2, 323, 324, 7, 113, 2, 2, 324, 325, 7, 112, 2, 2, 325,
	326, 7, 47, 2, 2, 326, 327, 7, 86, 2, 2, 327, 328, 7, 107, 2, 2, 328, 329,
	7, 118, 2, 2, 329, 330, 7, 110, 2, 2, 330, 372, 7, 103, 2, 2, 331, 332,
	7, 67, 2, 2, 332, 333, 7, 112, 2, 2, 333, 334, 7, 118, 2, 2, 334, 335,
	7, 47, 2, 2, 335, 336, 7, 88, 2, 2, 336, 337, 7, 103, 2, 2, 337, 338, 7,
	116, 2, 2, 338, 339, 7, 117, 2, 2, 339, 340, 7, 107, 2, 2, 340, 341, 7,
	113, 2, 2, 341, 372, 7, 112, 2, 2, 342, 343, 7, 85, 2, 2, 343, 344, 7,
	114, 2, 2, 344, 345, 7, 116, 2, 2, 345, 346, 7, 107, 2, 2, 346, 347, 7,
	112, 2, 2, 347, 348, 7, 105, 2, 2, 348, 349, 7, 47, 2, 2, 349, 350, 7,
	88, 2, 2, 350, 351, 7, 103, 2, 2, 351, 352, 7, 116, 2, 2, 352, 353, 7,
	117, 2, 2, 353, 354, 7, 107, 2, 2, 354, 355, 7, 113, 2, 2, 355, 372, 7,
	112, 2, 2, 356, 360, 5, 81, 41, 2, 357, 359, 5, 95, 48, 2, 358, 357, 3,
	2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2, 2,
	2, 361, 363, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 364, 7, 47, 2, 2, 364,
	368, 5, 81, 41, 2, 365, 367, 5, 95, 48, 2, 366, 365, 3, 2, 2, 2, 367, 370,
	3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 372, 3, 2,
	2, 2, 370, 368, 3, 2, 2, 2, 371, 109, 3, 2, 2, 2, 371, 125, 3, 2, 2, 2,
	371, 141, 3, 2, 2, 2, 371, 151, 3, 2, 2, 2, 371, 165, 3, 2, 2, 2, 371,
	179, 3, 2, 2, 2, 371, 193, 3, 2, 2, 2, 371, 204, 3, 2, 2, 2, 371, 222,
	3, 2, 2, 2, 371, 235, 3, 2, 2, 2, 371, 257, 3, 2, 2, 2, 371, 270, 3, 2,
	2, 2, 371, 289, 3, 2, 2, 2, 371, 311, 3, 2, 2, 2, 371, 331, 3, 2, 2, 2,
	371, 342, 3, 2, 2, 2, 371, 356, 3, 2, 2, 2, 372, 6, 3, 2, 2, 2, 373, 378,
	5, 9, 5, 2, 374, 377, 5, 9, 5, 2, 375, 377, 5, 85, 43, 2, 376, 374, 3,
	2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 380, 3, 2, 2, 2, 378, 376, 3, 2, 2,
	2, 378, 379, 3, 2, 2, 2, 379, 8, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 381,
	388, 9, 2, 2, 2, 382, 384, 5, 97, 49, 2, 383, 382, 3, 2, 2, 2, 384, 385,
	3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 388, 3, 2,
	2, 2, 387, 381, 3, 2, 2, 2, 387, 383, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2,
	389, 387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2, 391,
	393, 5, 11, 6, 2, 392, 387, 3, 2, 2, 2, 392, 391, 3, 2, 2, 2, 393, 10,
	3, 2, 2, 2, 394, 397, 5, 99, 50, 2, 395, 396, 7, 48, 2, 2, 396, 398, 5,
	97, 49, 2, 397, 395, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 397, 3, 2,
	2, 2, 399, 400, 3, 2, 2, 2, 400, 12, 3, 2, 2, 2, 401, 405, 5, 95, 48, 2,
	402, 404, 5, 97, 49, 2, 403, 402, 3, 2, 2, 2, 404, 407, 3, 2, 2, 2, 405,
	403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 14, 3, 2, 2, 2, 407, 405, 3,
	2, 2, 2, 408, 409, 7, 60, 2, 2, 409, 16, 3, 2, 2, 2, 410, 411, 7, 42, 2,
	2, 411, 18, 3, 2, 2, 2, 412, 413, 7, 43, 2, 2, 413, 20, 3, 2, 2, 2, 414,
	415, 7, 125, 2, 2, 415, 22, 3, 2, 2, 2, 416, 417, 7, 127, 2, 2, 417, 24,
	3, 2, 2, 2, 418, 419, 7, 93, 2, 2, 419, 26, 3, 2, 2, 2, 420, 421, 7, 95,
	2, 2, 421, 28, 3, 2, 2, 2, 422, 423, 7, 61, 2, 2, 423, 30, 3, 2, 2, 2,
	424, 425, 7, 46, 2, 2, 425, 32, 3, 2, 2, 2, 426, 427, 7, 48, 2, 2, 427,
	34, 3, 2, 2, 2, 428, 429, 7, 63, 2, 2, 429, 36, 3, 2, 2, 2, 430, 431, 7,
	64, 2, 2, 431, 38, 3, 2, 2, 2, 432, 433, 7, 62, 2, 2, 433, 40, 3, 2, 2,
	2, 434, 435, 7, 35, 2, 2, 435, 42, 3, 2, 2, 2, 436, 437, 7, 128, 2, 2,
	437, 44, 3, 2, 2, 2, 438, 439, 7, 65, 2, 2, 439, 46, 3, 2, 2, 2, 440, 441,
	7, 63, 2, 2, 441, 442, 7, 63, 2, 2, 442, 48, 3, 2, 2, 2, 443, 444, 7, 62,
	2, 2, 444, 445, 7, 63, 2, 2, 445, 50, 3, 2, 2, 2, 446, 447, 7, 64, 2, 2,
	447, 448, 7, 63, 2, 2, 448, 52, 3, 2, 2, 2, 449, 450, 7, 35, 2, 2, 450,
	451, 7, 63, 2, 2, 451, 54, 3, 2, 2, 2, 452, 453, 7, 40, 2, 2, 453, 454,
	7, 40, 2, 2, 454, 56, 3, 2, 2, 2, 455, 456, 7, 126, 2, 2, 456, 457, 7,
	126, 2, 2, 457, 58, 3, 2, 2, 2, 458, 459, 7, 45, 2, 2, 459, 460, 7, 45,
	2, 2, 460, 60, 3, 2, 2, 2, 461, 462, 7, 47, 2, 2, 462, 463, 7, 47, 2, 2,
	463, 62, 3, 2, 2, 2, 464, 465, 7, 45, 2, 2, 465, 64, 3, 2, 2, 2, 466, 467,
	7, 47, 2, 2, 467, 66, 3, 2, 2, 2, 468, 469, 7, 44, 2, 2, 469, 68, 3, 2,
	2, 2, 470, 471, 7, 49, 2, 2, 471, 70, 3, 2, 2, 2, 472, 473, 7, 40, 2, 2,
	473, 72, 3, 2, 2, 2, 474, 475, 7, 126, 2, 2, 475, 74, 3, 2, 2, 2, 476,
	477, 7, 96, 2, 2, 477, 76, 3, 2, 2, 2, 478, 479, 7, 39, 2, 2, 479, 78,
	3, 2, 2, 2, 480, 481, 7, 36, 2, 2, 481, 80, 3, 2, 2, 2, 482, 483, 9, 3,
	2, 2, 483, 82, 3, 2, 2, 2, 484, 488, 7, 61, 2, 2, 485, 487, 10, 4, 2, 2,
	486, 485, 3, 2, 2, 2, 487, 490, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2, 488,
	489, 3, 2, 2, 2, 489, 491, 3, 2, 2, 2, 490, 488, 3, 2, 2, 2, 491, 492,
	8, 42, 2, 2, 492, 84, 3, 2, 2, 2, 493, 494, 9, 5, 2, 2, 494, 86, 3, 2,
	2, 2, 495, 497, 7, 15, 2, 2, 496, 495, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2,
	497, 498, 3, 2, 2, 2, 498, 499, 7, 12, 2, 2, 499, 500, 3, 2, 2, 2, 500,
	501, 8, 44, 3, 2, 501, 88, 3, 2, 2, 2, 502, 507, 7, 36, 2, 2, 503, 506,
	10, 6, 2, 2, 504, 506, 5, 91, 46, 2, 505, 503, 3, 2, 2, 2, 505, 504, 3,
	2, 2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507, 508, 3, 2, 2,
	2, 508, 510, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 511, 7, 36, 2, 2, 511,
	90, 3, 2, 2, 2, 512, 513, 7, 94, 2, 2, 513, 534, 9, 7, 2, 2, 514, 519,
	7, 94, 2, 2, 515, 517, 9, 8, 2, 2, 516, 515, 3, 2, 2, 2, 516, 517, 3, 2,
	2, 2, 517, 518, 3, 2, 2, 2, 518, 520, 9, 9, 2, 2, 519, 516, 3, 2, 2, 2,
	519, 520, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 534, 9, 9, 2, 2, 522,
	524, 7, 94, 2, 2, 523, 525, 7, 119, 2, 2, 524, 523, 3, 2, 2, 2, 525, 526,
	3, 2, 2, 2, 526, 524, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 528, 3, 2,
	2, 2, 528, 529, 5, 93, 47, 2, 529, 530, 5, 93, 47, 2, 530, 531, 5, 93,
	47, 2, 531, 532, 5, 93, 47, 2, 532, 534, 3, 2, 2, 2, 533, 512, 3, 2, 2,
	2, 533, 514, 3, 2, 2, 2, 533, 522, 3, 2, 2, 2, 534, 92, 3, 2, 2, 2, 535,
	536, 9, 10, 2, 2, 536, 94, 3, 2, 2, 2, 537, 538, 9, 11, 2, 2, 538, 96,
	3, 2, 2, 2, 539, 542, 5, 95, 48, 2, 540, 542, 9, 12, 2, 2, 541, 539, 3,
	2, 2, 2, 541, 540, 3, 2, 2, 2, 542, 98, 3, 2, 2, 2, 543, 551, 9, 12, 2,
	2, 544, 546, 9, 13, 2, 2, 545, 544, 3, 2, 2, 2, 546, 549, 3, 2, 2, 2, 547,
	545, 3, 2, 2, 2, 547, 548, 3, 2, 2, 2, 548, 550, 3, 2, 2, 2, 549, 547,
	3, 2, 2, 2, 550, 552, 9, 12, 2, 2, 551, 547, 3, 2, 2, 2, 551, 552, 3, 2,
	2, 2, 552, 100, 3, 2, 2, 2, 25, 2, 360, 368, 371, 376, 378, 385, 387, 389,
	392, 399, 405, 488, 496, 505, 507, 516, 519, 526, 533, 541, 547, 551, 4,
	2, 3, 2, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'version'", "", "", "", "", "", "':'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'", "'~'", "'?'",
	"'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'", "'--'", "'+'",
	"'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "VERSION", "Key", "OTHERS", "ValueText", "Version", "IDENTIFIER", "COLON",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA",
	"DOT", "ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "DQUOTE", "Uppercase", "LINE_COMMENT",
	"SPACE", "NL", "STRING_LITERAL",
}

var lexerRuleNames = []string{
	"VERSION", "Key", "OTHERS", "ValueText", "Version", "IDENTIFIER", "COLON",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA",
	"DOT", "ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "DQUOTE", "Uppercase", "LINE_COMMENT",
	"SPACE", "NL", "STRING_LITERAL", "EscapeSequence", "HexDigit", "Letter",
	"LetterOrDigit", "Digits",
}

type ManifestLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewManifestLexer(input antlr.CharStream) *ManifestLexer {

	l := new(ManifestLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Manifest.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ManifestLexer tokens.
const (
	ManifestLexerVERSION        = 1
	ManifestLexerKey            = 2
	ManifestLexerOTHERS         = 3
	ManifestLexerValueText      = 4
	ManifestLexerVersion        = 5
	ManifestLexerIDENTIFIER     = 6
	ManifestLexerCOLON          = 7
	ManifestLexerLPAREN         = 8
	ManifestLexerRPAREN         = 9
	ManifestLexerLBRACE         = 10
	ManifestLexerRBRACE         = 11
	ManifestLexerLBRACK         = 12
	ManifestLexerRBRACK         = 13
	ManifestLexerSEMI           = 14
	ManifestLexerCOMMA          = 15
	ManifestLexerDOT            = 16
	ManifestLexerASSIGN         = 17
	ManifestLexerGT             = 18
	ManifestLexerLT             = 19
	ManifestLexerBANG           = 20
	ManifestLexerTILDE          = 21
	ManifestLexerQUESTION       = 22
	ManifestLexerEQUAL          = 23
	ManifestLexerLE             = 24
	ManifestLexerGE             = 25
	ManifestLexerNOTEQUAL       = 26
	ManifestLexerAND            = 27
	ManifestLexerOR             = 28
	ManifestLexerINC            = 29
	ManifestLexerDEC            = 30
	ManifestLexerADD            = 31
	ManifestLexerSUB            = 32
	ManifestLexerMUL            = 33
	ManifestLexerDIV            = 34
	ManifestLexerBITAND         = 35
	ManifestLexerBITOR          = 36
	ManifestLexerCARET          = 37
	ManifestLexerMOD            = 38
	ManifestLexerDQUOTE         = 39
	ManifestLexerUppercase      = 40
	ManifestLexerLINE_COMMENT   = 41
	ManifestLexerSPACE          = 42
	ManifestLexerNL             = 43
	ManifestLexerSTRING_LITERAL = 44
)
