// Code generated from ./src/lib/parser/httpSpec.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // httpSpec

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 334,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 7, 2, 80, 10, 2, 12, 2, 14, 2, 83, 11, 2, 3, 2, 3, 2, 7, 2, 87,
	10, 2, 12, 2, 14, 2, 90, 11, 2, 3, 2, 7, 2, 93, 10, 2, 12, 2, 14, 2, 96,
	11, 2, 3, 3, 6, 3, 99, 10, 3, 13, 3, 14, 3, 100, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 110, 10, 4, 3, 4, 5, 4, 113, 10, 4, 3, 4, 5, 4,
	116, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 121, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	126, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 132, 10, 7, 13, 7, 14, 7, 133,
	3, 7, 3, 7, 6, 7, 138, 10, 7, 13, 7, 14, 7, 139, 3, 8, 3, 8, 3, 8, 5, 8,
	145, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 150, 10, 9, 3, 9, 3, 9, 5, 9, 154,
	10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 159, 10, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 164, 10, 10, 3, 10, 3, 10, 5, 10, 168, 10, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 5, 12, 174, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 181,
	10, 14, 3, 15, 6, 15, 184, 10, 15, 13, 15, 14, 15, 185, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 193, 10, 16, 3, 17, 6, 17, 196, 10, 17, 13, 17,
	14, 17, 197, 3, 18, 6, 18, 201, 10, 18, 13, 18, 14, 18, 202, 3, 19, 3,
	19, 3, 19, 3, 19, 6, 19, 209, 10, 19, 13, 19, 14, 19, 210, 5, 19, 213,
	10, 19, 3, 20, 3, 20, 3, 21, 7, 21, 218, 10, 21, 12, 21, 14, 21, 221, 11,
	21, 3, 22, 7, 22, 224, 10, 22, 12, 22, 14, 22, 227, 11, 22, 3, 22, 3, 22,
	5, 22, 231, 10, 22, 3, 23, 7, 23, 234, 10, 23, 12, 23, 14, 23, 237, 11,
	23, 3, 23, 3, 23, 5, 23, 241, 10, 23, 3, 24, 3, 24, 3, 24, 7, 24, 246,
	10, 24, 12, 24, 14, 24, 249, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 26, 6, 26, 258, 10, 26, 13, 26, 14, 26, 259, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 5, 28, 268, 10, 28, 3, 29, 3, 29, 3, 29, 5, 29,
	273, 10, 29, 3, 30, 3, 30, 3, 30, 5, 30, 278, 10, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 5, 33, 288, 10, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 296, 10, 34, 12, 34, 14, 34, 299, 11,
	34, 3, 34, 3, 34, 5, 34, 303, 10, 34, 3, 35, 6, 35, 306, 10, 35, 13, 35,
	14, 35, 307, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 317,
	10, 36, 3, 37, 6, 37, 320, 10, 37, 13, 37, 14, 37, 321, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 2, 2, 40, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	2, 13, 3, 2, 3, 11, 3, 2, 17, 18, 3, 2, 22, 23, 5, 2, 14, 15, 20, 20, 23,
	23, 4, 2, 23, 23, 41, 41, 4, 2, 14, 15, 23, 23, 3, 2, 15, 15, 3, 2, 14,
	14, 3, 2, 20, 20, 3, 2, 24, 26, 7, 2, 13, 14, 20, 20, 23, 23, 27, 34, 47,
	48, 2, 337, 2, 81, 3, 2, 2, 2, 4, 98, 3, 2, 2, 2, 6, 104, 3, 2, 2, 2, 8,
	120, 3, 2, 2, 2, 10, 127, 3, 2, 2, 2, 12, 129, 3, 2, 2, 2, 14, 144, 3,
	2, 2, 2, 16, 146, 3, 2, 2, 2, 18, 158, 3, 2, 2, 2, 20, 169, 3, 2, 2, 2,
	22, 171, 3, 2, 2, 2, 24, 175, 3, 2, 2, 2, 26, 177, 3, 2, 2, 2, 28, 183,
	3, 2, 2, 2, 30, 192, 3, 2, 2, 2, 32, 195, 3, 2, 2, 2, 34, 200, 3, 2, 2,
	2, 36, 212, 3, 2, 2, 2, 38, 214, 3, 2, 2, 2, 40, 219, 3, 2, 2, 2, 42, 225,
	3, 2, 2, 2, 44, 235, 3, 2, 2, 2, 46, 247, 3, 2, 2, 2, 48, 250, 3, 2, 2,
	2, 50, 257, 3, 2, 2, 2, 52, 261, 3, 2, 2, 2, 54, 267, 3, 2, 2, 2, 56, 269,
	3, 2, 2, 2, 58, 277, 3, 2, 2, 2, 60, 279, 3, 2, 2, 2, 62, 283, 3, 2, 2,
	2, 64, 285, 3, 2, 2, 2, 66, 291, 3, 2, 2, 2, 68, 305, 3, 2, 2, 2, 70, 309,
	3, 2, 2, 2, 72, 319, 3, 2, 2, 2, 74, 323, 3, 2, 2, 2, 76, 327, 3, 2, 2,
	2, 78, 80, 7, 51, 2, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	84, 88, 5, 6, 4, 2, 85, 87, 5, 4, 3, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 94, 3, 2, 2, 2, 90,
	88, 3, 2, 2, 2, 91, 93, 7, 51, 2, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2,
	2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 3, 3, 2, 2, 2, 96, 94,
	3, 2, 2, 2, 97, 99, 7, 51, 2, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2,
	2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102,
	103, 5, 6, 4, 2, 103, 5, 3, 2, 2, 2, 104, 105, 5, 8, 5, 2, 105, 106, 7,
	40, 2, 2, 106, 107, 5, 46, 24, 2, 107, 109, 7, 40, 2, 2, 108, 110, 5, 54,
	28, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2,
	111, 113, 5, 70, 36, 2, 112, 111, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113,
	115, 3, 2, 2, 2, 114, 116, 5, 74, 38, 2, 115, 114, 3, 2, 2, 2, 115, 116,
	3, 2, 2, 2, 116, 7, 3, 2, 2, 2, 117, 118, 5, 10, 6, 2, 118, 119, 7, 45,
	2, 2, 119, 121, 3, 2, 2, 2, 120, 117, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2,
	121, 122, 3, 2, 2, 2, 122, 125, 5, 14, 8, 2, 123, 124, 7, 45, 2, 2, 124,
	126, 5, 12, 7, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 9, 3,
	2, 2, 2, 127, 128, 9, 2, 2, 2, 128, 11, 3, 2, 2, 2, 129, 131, 7, 12, 2,
	2, 130, 132, 7, 48, 2, 2, 131, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137,
	7, 13, 2, 2, 136, 138, 7, 48, 2, 2, 137, 136, 3, 2, 2, 2, 138, 139, 3,
	2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 13, 3, 2, 2,
	2, 141, 145, 5, 16, 9, 2, 142, 145, 5, 18, 10, 2, 143, 145, 5, 24, 13,
	2, 144, 141, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145,
	15, 3, 2, 2, 2, 146, 149, 5, 36, 19, 2, 147, 148, 7, 14, 2, 2, 148, 150,
	5, 42, 22, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 153, 3,
	2, 2, 2, 151, 152, 7, 15, 2, 2, 152, 154, 5, 44, 23, 2, 153, 151, 3, 2,
	2, 2, 153, 154, 3, 2, 2, 2, 154, 17, 3, 2, 2, 2, 155, 156, 5, 20, 11, 2,
	156, 157, 7, 16, 2, 2, 157, 159, 3, 2, 2, 2, 158, 155, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 163, 5, 22, 12, 2, 161, 162,
	7, 14, 2, 2, 162, 164, 5, 42, 22, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3,
	2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 166, 7, 15, 2, 2, 166, 168, 5, 44,
	23, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 19, 3, 2, 2, 2,
	169, 170, 9, 3, 2, 2, 170, 21, 3, 2, 2, 2, 171, 173, 5, 26, 14, 2, 172,
	174, 5, 36, 19, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 23,
	3, 2, 2, 2, 175, 176, 7, 19, 2, 2, 176, 25, 3, 2, 2, 2, 177, 180, 5, 30,
	16, 2, 178, 179, 7, 20, 2, 2, 179, 181, 5, 28, 15, 2, 180, 178, 3, 2, 2,
	2, 180, 181, 3, 2, 2, 2, 181, 27, 3, 2, 2, 2, 182, 184, 7, 48, 2, 2, 183,
	182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186,
	3, 2, 2, 2, 186, 29, 3, 2, 2, 2, 187, 188, 7, 21, 2, 2, 188, 189, 5, 32,
	17, 2, 189, 190, 7, 22, 2, 2, 190, 193, 3, 2, 2, 2, 191, 193, 5, 34, 18,
	2, 192, 187, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 31, 3, 2, 2, 2, 194,
	196, 10, 4, 2, 2, 195, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 195,
	3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 33, 3, 2, 2, 2, 199, 201, 10, 5,
	2, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2,
	202, 203, 3, 2, 2, 2, 203, 35, 3, 2, 2, 2, 204, 213, 7, 23, 2, 2, 205,
	206, 5, 38, 20, 2, 206, 207, 5, 40, 21, 2, 207, 209, 3, 2, 2, 2, 208, 205,
	3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2,
	2, 2, 211, 213, 3, 2, 2, 2, 212, 204, 3, 2, 2, 2, 212, 208, 3, 2, 2, 2,
	213, 37, 3, 2, 2, 2, 214, 215, 9, 6, 2, 2, 215, 39, 3, 2, 2, 2, 216, 218,
	10, 7, 2, 2, 217, 216, 3, 2, 2, 2, 218, 221, 3, 2, 2, 2, 219, 217, 3, 2,
	2, 2, 219, 220, 3, 2, 2, 2, 220, 41, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2,
	222, 224, 10, 8, 2, 2, 223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 230, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 228, 229, 7, 41, 2, 2, 229, 231, 5, 42, 22, 2, 230, 228, 3,
	2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 43, 3, 2, 2, 2, 232, 234, 10, 9, 2,
	2, 233, 232, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235,
	236, 3, 2, 2, 2, 236, 240, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 239,
	7, 41, 2, 2, 239, 241, 5, 44, 23, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3,
	2, 2, 2, 241, 45, 3, 2, 2, 2, 242, 243, 5, 48, 25, 2, 243, 244, 7, 40,
	2, 2, 244, 246, 3, 2, 2, 2, 245, 242, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2,
	247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 47, 3, 2, 2, 2, 249, 247,
	3, 2, 2, 2, 250, 251, 5, 50, 26, 2, 251, 252, 7, 20, 2, 2, 252, 253, 7,
	44, 2, 2, 253, 254, 5, 52, 27, 2, 254, 255, 7, 44, 2, 2, 255, 49, 3, 2,
	2, 2, 256, 258, 10, 10, 2, 2, 257, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2,
	2, 259, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 51, 3, 2, 2, 2, 261,
	262, 7, 42, 2, 2, 262, 263, 7, 41, 2, 2, 263, 264, 5, 52, 27, 2, 264, 53,
	3, 2, 2, 2, 265, 268, 5, 56, 29, 2, 266, 268, 5, 64, 33, 2, 267, 265, 3,
	2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 55, 3, 2, 2, 2, 269, 272, 5, 58, 30,
	2, 270, 271, 7, 40, 2, 2, 271, 273, 5, 58, 30, 2, 272, 270, 3, 2, 2, 2,
	272, 273, 3, 2, 2, 2, 273, 57, 3, 2, 2, 2, 274, 275, 10, 11, 2, 2, 275,
	278, 7, 42, 2, 2, 276, 278, 5, 60, 31, 2, 277, 274, 3, 2, 2, 2, 277, 276,
	3, 2, 2, 2, 278, 59, 3, 2, 2, 2, 279, 280, 7, 24, 2, 2, 280, 281, 7, 45,
	2, 2, 281, 282, 5, 62, 32, 2, 282, 61, 3, 2, 2, 2, 283, 284, 7, 42, 2,
	2, 284, 63, 3, 2, 2, 2, 285, 287, 5, 66, 34, 2, 286, 288, 5, 64, 33, 2,
	287, 286, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289,
	290, 5, 68, 35, 2, 290, 65, 3, 2, 2, 2, 291, 297, 5, 68, 35, 2, 292, 293,
	5, 48, 25, 2, 293, 294, 7, 40, 2, 2, 294, 296, 3, 2, 2, 2, 295, 292, 3,
	2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2,
	2, 298, 300, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 302, 7, 40, 2, 2, 301,
	303, 5, 56, 29, 2, 302, 301, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 67,
	3, 2, 2, 2, 304, 306, 9, 12, 2, 2, 305, 304, 3, 2, 2, 2, 306, 307, 3, 2,
	2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 69, 3, 2, 2, 2,
	309, 310, 7, 35, 2, 2, 310, 316, 7, 45, 2, 2, 311, 312, 7, 36, 2, 2, 312,
	313, 5, 72, 37, 2, 313, 314, 7, 37, 2, 2, 314, 317, 3, 2, 2, 2, 315, 317,
	5, 62, 32, 2, 316, 311, 3, 2, 2, 2, 316, 315, 3, 2, 2, 2, 317, 71, 3, 2,
	2, 2, 318, 320, 7, 46, 2, 2, 319, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2,
	321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 73, 3, 2, 2, 2, 323, 324,
	7, 25, 2, 2, 324, 325, 7, 45, 2, 2, 325, 326, 5, 62, 32, 2, 326, 75, 3,
	2, 2, 2, 327, 328, 7, 38, 2, 2, 328, 329, 7, 44, 2, 2, 329, 330, 7, 49,
	2, 2, 330, 331, 7, 44, 2, 2, 331, 332, 7, 39, 2, 2, 332, 77, 3, 2, 2, 2,
	43, 81, 88, 94, 100, 109, 112, 115, 120, 125, 133, 139, 144, 149, 153,
	158, 163, 167, 173, 180, 185, 192, 197, 202, 210, 212, 219, 225, 230, 235,
	240, 247, 259, 267, 272, 277, 287, 297, 302, 307, 316, 321,
}
var literalNames = []string{
	"", "'GET'", "'HEAD'", "'POST'", "'PUT'", "'DELETE'", "'CONNECT'", "'PATCH'",
	"'OPTIONS'", "'TRACE'", "'HTTP/'", "'.'", "'?'", "'#'", "'//'", "'http'",
	"'https'", "'*'", "':'", "'['", "']'", "'/'", "'<'", "'<>'", "'###'", "'''",
	"'('", "')'", "'+'", "'_'", "','", "'-'", "'='", "'>'", "'{%'", "'%}'",
	"'{{'", "'}}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "WHITESPACE", "OPTIONALWHITESPACE",
	"REQUIREDWHITESPACE", "INPUT", "ALPHA", "DIGIT", "ID", "LINECOMMENT", "REQUESTSEPARATOR",
}

var ruleNames = []string{
	"requestfile", "requestwithseparator", "request", "requestline", "method",
	"httpversion", "requesttarget", "originform", "absoluteform", "scheme",
	"hierpart", "asteriskform", "authority", "port", "host", "ipv6address",
	"ipv4addressorregname", "absolutepath", "pathseparator", "segment", "query",
	"requestfragment", "headers", "headerfield", "fieldname", "fieldvalue",
	"messagebody", "messages", "messageline", "inputfileref", "filepath", "multipartformdata",
	"multipartfield", "boundary", "responsehandler", "handlerscript", "responseref",
	"envvariable",
}

type httpSpecParser struct {
	*antlr.BaseParser
}

// NewhttpSpecParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *httpSpecParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpSpecParser(input antlr.TokenStream) *httpSpecParser {
	this := new(httpSpecParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "httpSpec.g4"

	return this
}

// httpSpecParser tokens.
const (
	httpSpecParserEOF                = antlr.TokenEOF
	httpSpecParserT__0               = 1
	httpSpecParserT__1               = 2
	httpSpecParserT__2               = 3
	httpSpecParserT__3               = 4
	httpSpecParserT__4               = 5
	httpSpecParserT__5               = 6
	httpSpecParserT__6               = 7
	httpSpecParserT__7               = 8
	httpSpecParserT__8               = 9
	httpSpecParserT__9               = 10
	httpSpecParserT__10              = 11
	httpSpecParserT__11              = 12
	httpSpecParserT__12              = 13
	httpSpecParserT__13              = 14
	httpSpecParserT__14              = 15
	httpSpecParserT__15              = 16
	httpSpecParserT__16              = 17
	httpSpecParserT__17              = 18
	httpSpecParserT__18              = 19
	httpSpecParserT__19              = 20
	httpSpecParserT__20              = 21
	httpSpecParserT__21              = 22
	httpSpecParserT__22              = 23
	httpSpecParserT__23              = 24
	httpSpecParserT__24              = 25
	httpSpecParserT__25              = 26
	httpSpecParserT__26              = 27
	httpSpecParserT__27              = 28
	httpSpecParserT__28              = 29
	httpSpecParserT__29              = 30
	httpSpecParserT__30              = 31
	httpSpecParserT__31              = 32
	httpSpecParserT__32              = 33
	httpSpecParserT__33              = 34
	httpSpecParserT__34              = 35
	httpSpecParserT__35              = 36
	httpSpecParserT__36              = 37
	httpSpecParserNEWLINE            = 38
	httpSpecParserNEWLINEWITHINDENT  = 39
	httpSpecParserLINETAIL           = 40
	httpSpecParserWHITESPACE         = 41
	httpSpecParserOPTIONALWHITESPACE = 42
	httpSpecParserREQUIREDWHITESPACE = 43
	httpSpecParserINPUT              = 44
	httpSpecParserALPHA              = 45
	httpSpecParserDIGIT              = 46
	httpSpecParserID                 = 47
	httpSpecParserLINECOMMENT        = 48
	httpSpecParserREQUESTSEPARATOR   = 49
)

// httpSpecParser rules.
const (
	httpSpecParserRULE_requestfile          = 0
	httpSpecParserRULE_requestwithseparator = 1
	httpSpecParserRULE_request              = 2
	httpSpecParserRULE_requestline          = 3
	httpSpecParserRULE_method               = 4
	httpSpecParserRULE_httpversion          = 5
	httpSpecParserRULE_requesttarget        = 6
	httpSpecParserRULE_originform           = 7
	httpSpecParserRULE_absoluteform         = 8
	httpSpecParserRULE_scheme               = 9
	httpSpecParserRULE_hierpart             = 10
	httpSpecParserRULE_asteriskform         = 11
	httpSpecParserRULE_authority            = 12
	httpSpecParserRULE_port                 = 13
	httpSpecParserRULE_host                 = 14
	httpSpecParserRULE_ipv6address          = 15
	httpSpecParserRULE_ipv4addressorregname = 16
	httpSpecParserRULE_absolutepath         = 17
	httpSpecParserRULE_pathseparator        = 18
	httpSpecParserRULE_segment              = 19
	httpSpecParserRULE_query                = 20
	httpSpecParserRULE_requestfragment      = 21
	httpSpecParserRULE_headers              = 22
	httpSpecParserRULE_headerfield          = 23
	httpSpecParserRULE_fieldname            = 24
	httpSpecParserRULE_fieldvalue           = 25
	httpSpecParserRULE_messagebody          = 26
	httpSpecParserRULE_messages             = 27
	httpSpecParserRULE_messageline          = 28
	httpSpecParserRULE_inputfileref         = 29
	httpSpecParserRULE_filepath             = 30
	httpSpecParserRULE_multipartformdata    = 31
	httpSpecParserRULE_multipartfield       = 32
	httpSpecParserRULE_boundary             = 33
	httpSpecParserRULE_responsehandler      = 34
	httpSpecParserRULE_handlerscript        = 35
	httpSpecParserRULE_responseref          = 36
	httpSpecParserRULE_envvariable          = 37
)

// IRequestfileContext is an interface to support dynamic dispatch.
type IRequestfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestfileContext differentiates from other interfaces.
	IsRequestfileContext()
}

type RequestfileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestfileContext() *RequestfileContext {
	var p = new(RequestfileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_requestfile
	return p
}

func (*RequestfileContext) IsRequestfileContext() {}

func NewRequestfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestfileContext {
	var p = new(RequestfileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestfile

	return p
}

func (s *RequestfileContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestfileContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RequestfileContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserREQUESTSEPARATOR)
}

func (s *RequestfileContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTSEPARATOR, i)
}

func (s *RequestfileContext) AllRequestwithseparator() []IRequestwithseparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRequestwithseparatorContext)(nil)).Elem())
	var tst = make([]IRequestwithseparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRequestwithseparatorContext)
		}
	}

	return tst
}

func (s *RequestfileContext) Requestwithseparator(i int) IRequestwithseparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestwithseparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRequestwithseparatorContext)
}

func (s *RequestfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestfileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestfile(s)
	}
}

func (s *RequestfileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestfile(s)
	}
}

func (p *httpSpecParser) Requestfile() (localctx IRequestfileContext) {
	localctx = NewRequestfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, httpSpecParserRULE_requestfile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(76)
				p.Match(httpSpecParserREQUESTSEPARATOR)
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(82)
		p.Request()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(83)
				p.Requestwithseparator()
			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpSpecParserREQUESTSEPARATOR {
		{
			p.SetState(89)
			p.Match(httpSpecParserREQUESTSEPARATOR)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRequestwithseparatorContext is an interface to support dynamic dispatch.
type IRequestwithseparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestwithseparatorContext differentiates from other interfaces.
	IsRequestwithseparatorContext()
}

type RequestwithseparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestwithseparatorContext() *RequestwithseparatorContext {
	var p = new(RequestwithseparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_requestwithseparator
	return p
}

func (*RequestwithseparatorContext) IsRequestwithseparatorContext() {}

func NewRequestwithseparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestwithseparatorContext {
	var p = new(RequestwithseparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestwithseparator

	return p
}

func (s *RequestwithseparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestwithseparatorContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *RequestwithseparatorContext) AllREQUESTSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserREQUESTSEPARATOR)
}

func (s *RequestwithseparatorContext) REQUESTSEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUESTSEPARATOR, i)
}

func (s *RequestwithseparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestwithseparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestwithseparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestwithseparator(s)
	}
}

func (s *RequestwithseparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestwithseparator(s)
	}
}

func (p *httpSpecParser) Requestwithseparator() (localctx IRequestwithseparatorContext) {
	localctx = NewRequestwithseparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, httpSpecParserRULE_requestwithseparator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(95)
				p.Match(httpSpecParserREQUESTSEPARATOR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(100)
		p.Request()
	}

	return localctx
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) Requestline() IRequestlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestlineContext)
}

func (s *RequestContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserNEWLINE)
}

func (s *RequestContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, i)
}

func (s *RequestContext) Headers() IHeadersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeadersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeadersContext)
}

func (s *RequestContext) Messagebody() IMessagebodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagebodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagebodyContext)
}

func (s *RequestContext) Responsehandler() IResponsehandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResponsehandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResponsehandlerContext)
}

func (s *RequestContext) Responseref() IResponserefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResponserefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResponserefContext)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *httpSpecParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, httpSpecParserRULE_request)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Requestline()
	}
	{
		p.SetState(103)
		p.Match(httpSpecParserNEWLINE)
	}
	{
		p.SetState(104)
		p.Headers()
	}
	{
		p.SetState(105)
		p.Match(httpSpecParserNEWLINE)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(106)
			p.Messagebody()
		}

	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__32 {
		{
			p.SetState(109)
			p.Responsehandler()
		}

	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__22 {
		{
			p.SetState(112)
			p.Responseref()
		}

	}

	return localctx
}

// IRequestlineContext is an interface to support dynamic dispatch.
type IRequestlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestlineContext differentiates from other interfaces.
	IsRequestlineContext()
}

type RequestlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestlineContext() *RequestlineContext {
	var p = new(RequestlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_requestline
	return p
}

func (*RequestlineContext) IsRequestlineContext() {}

func NewRequestlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestlineContext {
	var p = new(RequestlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestline

	return p
}

func (s *RequestlineContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestlineContext) Requesttarget() IRequesttargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequesttargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequesttargetContext)
}

func (s *RequestlineContext) Method() IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *RequestlineContext) AllREQUIREDWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserREQUIREDWHITESPACE)
}

func (s *RequestlineContext) REQUIREDWHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUIREDWHITESPACE, i)
}

func (s *RequestlineContext) Httpversion() IHttpversionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHttpversionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHttpversionContext)
}

func (s *RequestlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestline(s)
	}
}

func (s *RequestlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestline(s)
	}
}

func (p *httpSpecParser) Requestline() (localctx IRequestlineContext) {
	localctx = NewRequestlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, httpSpecParserRULE_requestline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(115)
			p.Method()
		}
		{
			p.SetState(116)
			p.Match(httpSpecParserREQUIREDWHITESPACE)
		}

	}
	{
		p.SetState(120)
		p.Requesttarget()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserREQUIREDWHITESPACE {
		{
			p.SetState(121)
			p.Match(httpSpecParserREQUIREDWHITESPACE)
		}
		{
			p.SetState(122)
			p.Httpversion()
		}

	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }
func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *httpSpecParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, httpSpecParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__0)|(1<<httpSpecParserT__1)|(1<<httpSpecParserT__2)|(1<<httpSpecParserT__3)|(1<<httpSpecParserT__4)|(1<<httpSpecParserT__5)|(1<<httpSpecParserT__6)|(1<<httpSpecParserT__7)|(1<<httpSpecParserT__8))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHttpversionContext is an interface to support dynamic dispatch.
type IHttpversionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHttpversionContext differentiates from other interfaces.
	IsHttpversionContext()
}

type HttpversionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpversionContext() *HttpversionContext {
	var p = new(HttpversionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_httpversion
	return p
}

func (*HttpversionContext) IsHttpversionContext() {}

func NewHttpversionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpversionContext {
	var p = new(HttpversionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_httpversion

	return p
}

func (s *HttpversionContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpversionContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserDIGIT)
}

func (s *HttpversionContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserDIGIT, i)
}

func (s *HttpversionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpversionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpversionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHttpversion(s)
	}
}

func (s *HttpversionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHttpversion(s)
	}
}

func (p *httpSpecParser) Httpversion() (localctx IHttpversionContext) {
	localctx = NewHttpversionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, httpSpecParserRULE_httpversion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(httpSpecParserT__9)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserDIGIT {
		{
			p.SetState(128)
			p.Match(httpSpecParserDIGIT)
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(httpSpecParserT__10)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserDIGIT {
		{
			p.SetState(134)
			p.Match(httpSpecParserDIGIT)
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRequesttargetContext is an interface to support dynamic dispatch.
type IRequesttargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequesttargetContext differentiates from other interfaces.
	IsRequesttargetContext()
}

type RequesttargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequesttargetContext() *RequesttargetContext {
	var p = new(RequesttargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_requesttarget
	return p
}

func (*RequesttargetContext) IsRequesttargetContext() {}

func NewRequesttargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequesttargetContext {
	var p = new(RequesttargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requesttarget

	return p
}

func (s *RequesttargetContext) GetParser() antlr.Parser { return s.parser }

func (s *RequesttargetContext) Originform() IOriginformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOriginformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOriginformContext)
}

func (s *RequesttargetContext) Absoluteform() IAbsoluteformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsoluteformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsoluteformContext)
}

func (s *RequesttargetContext) Asteriskform() IAsteriskformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsteriskformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsteriskformContext)
}

func (s *RequesttargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequesttargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequesttargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequesttarget(s)
	}
}

func (s *RequesttargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequesttarget(s)
	}
}

func (p *httpSpecParser) Requesttarget() (localctx IRequesttargetContext) {
	localctx = NewRequesttargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, httpSpecParserRULE_requesttarget)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Originform()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Absoluteform()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Asteriskform()
		}

	}

	return localctx
}

// IOriginformContext is an interface to support dynamic dispatch.
type IOriginformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOriginformContext differentiates from other interfaces.
	IsOriginformContext()
}

type OriginformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOriginformContext() *OriginformContext {
	var p = new(OriginformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_originform
	return p
}

func (*OriginformContext) IsOriginformContext() {}

func NewOriginformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OriginformContext {
	var p = new(OriginformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_originform

	return p
}

func (s *OriginformContext) GetParser() antlr.Parser { return s.parser }

func (s *OriginformContext) Absolutepath() IAbsolutepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsolutepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsolutepathContext)
}

func (s *OriginformContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *OriginformContext) Requestfragment() IRequestfragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestfragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestfragmentContext)
}

func (s *OriginformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OriginformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OriginformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterOriginform(s)
	}
}

func (s *OriginformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitOriginform(s)
	}
}

func (p *httpSpecParser) Originform() (localctx IOriginformContext) {
	localctx = NewOriginformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, httpSpecParserRULE_originform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Absolutepath()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__11 {
		{
			p.SetState(145)
			p.Match(httpSpecParserT__11)
		}
		{
			p.SetState(146)
			p.Query()
		}

	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__12 {
		{
			p.SetState(149)
			p.Match(httpSpecParserT__12)
		}
		{
			p.SetState(150)
			p.Requestfragment()
		}

	}

	return localctx
}

// IAbsoluteformContext is an interface to support dynamic dispatch.
type IAbsoluteformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsoluteformContext differentiates from other interfaces.
	IsAbsoluteformContext()
}

type AbsoluteformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsoluteformContext() *AbsoluteformContext {
	var p = new(AbsoluteformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_absoluteform
	return p
}

func (*AbsoluteformContext) IsAbsoluteformContext() {}

func NewAbsoluteformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsoluteformContext {
	var p = new(AbsoluteformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_absoluteform

	return p
}

func (s *AbsoluteformContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsoluteformContext) Hierpart() IHierpartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHierpartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHierpartContext)
}

func (s *AbsoluteformContext) Scheme() ISchemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemeContext)
}

func (s *AbsoluteformContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *AbsoluteformContext) Requestfragment() IRequestfragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestfragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestfragmentContext)
}

func (s *AbsoluteformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsoluteformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsoluteformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterAbsoluteform(s)
	}
}

func (s *AbsoluteformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitAbsoluteform(s)
	}
}

func (p *httpSpecParser) Absoluteform() (localctx IAbsoluteformContext) {
	localctx = NewAbsoluteformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, httpSpecParserRULE_absoluteform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.Scheme()
		}
		{
			p.SetState(154)
			p.Match(httpSpecParserT__13)
		}

	}
	{
		p.SetState(158)
		p.Hierpart()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__11 {
		{
			p.SetState(159)
			p.Match(httpSpecParserT__11)
		}
		{
			p.SetState(160)
			p.Query()
		}

	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__12 {
		{
			p.SetState(163)
			p.Match(httpSpecParserT__12)
		}
		{
			p.SetState(164)
			p.Requestfragment()
		}

	}

	return localctx
}

// ISchemeContext is an interface to support dynamic dispatch.
type ISchemeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemeContext differentiates from other interfaces.
	IsSchemeContext()
}

type SchemeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemeContext() *SchemeContext {
	var p = new(SchemeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_scheme
	return p
}

func (*SchemeContext) IsSchemeContext() {}

func NewSchemeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemeContext {
	var p = new(SchemeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_scheme

	return p
}

func (s *SchemeContext) GetParser() antlr.Parser { return s.parser }
func (s *SchemeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterScheme(s)
	}
}

func (s *SchemeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitScheme(s)
	}
}

func (p *httpSpecParser) Scheme() (localctx ISchemeContext) {
	localctx = NewSchemeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, httpSpecParserRULE_scheme)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(_la == httpSpecParserT__14 || _la == httpSpecParserT__15) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHierpartContext is an interface to support dynamic dispatch.
type IHierpartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHierpartContext differentiates from other interfaces.
	IsHierpartContext()
}

type HierpartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHierpartContext() *HierpartContext {
	var p = new(HierpartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_hierpart
	return p
}

func (*HierpartContext) IsHierpartContext() {}

func NewHierpartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HierpartContext {
	var p = new(HierpartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_hierpart

	return p
}

func (s *HierpartContext) GetParser() antlr.Parser { return s.parser }

func (s *HierpartContext) Authority() IAuthorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuthorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *HierpartContext) Absolutepath() IAbsolutepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsolutepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsolutepathContext)
}

func (s *HierpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HierpartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HierpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHierpart(s)
	}
}

func (s *HierpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHierpart(s)
	}
}

func (p *httpSpecParser) Hierpart() (localctx IHierpartContext) {
	localctx = NewHierpartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, httpSpecParserRULE_hierpart)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Authority()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__20 || _la == httpSpecParserNEWLINEWITHINDENT {
		{
			p.SetState(170)
			p.Absolutepath()
		}

	}

	return localctx
}

// IAsteriskformContext is an interface to support dynamic dispatch.
type IAsteriskformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsteriskformContext differentiates from other interfaces.
	IsAsteriskformContext()
}

type AsteriskformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsteriskformContext() *AsteriskformContext {
	var p = new(AsteriskformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_asteriskform
	return p
}

func (*AsteriskformContext) IsAsteriskformContext() {}

func NewAsteriskformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsteriskformContext {
	var p = new(AsteriskformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_asteriskform

	return p
}

func (s *AsteriskformContext) GetParser() antlr.Parser { return s.parser }
func (s *AsteriskformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsteriskformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsteriskformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterAsteriskform(s)
	}
}

func (s *AsteriskformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitAsteriskform(s)
	}
}

func (p *httpSpecParser) Asteriskform() (localctx IAsteriskformContext) {
	localctx = NewAsteriskformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, httpSpecParserRULE_asteriskform)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(httpSpecParserT__16)
	}

	return localctx
}

// IAuthorityContext is an interface to support dynamic dispatch.
type IAuthorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuthorityContext differentiates from other interfaces.
	IsAuthorityContext()
}

type AuthorityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorityContext() *AuthorityContext {
	var p = new(AuthorityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_authority
	return p
}

func (*AuthorityContext) IsAuthorityContext() {}

func NewAuthorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorityContext {
	var p = new(AuthorityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_authority

	return p
}

func (s *AuthorityContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorityContext) Host() IHostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostContext)
}

func (s *AuthorityContext) Port() IPortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *AuthorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthorityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterAuthority(s)
	}
}

func (s *AuthorityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitAuthority(s)
	}
}

func (p *httpSpecParser) Authority() (localctx IAuthorityContext) {
	localctx = NewAuthorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, httpSpecParserRULE_authority)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Host()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserT__17 {
		{
			p.SetState(176)
			p.Match(httpSpecParserT__17)
		}
		{
			p.SetState(177)
			p.Port()
		}

	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserDIGIT)
}

func (s *PortContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserDIGIT, i)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitPort(s)
	}
}

func (p *httpSpecParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, httpSpecParserRULE_port)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserDIGIT {
		{
			p.SetState(180)
			p.Match(httpSpecParserDIGIT)
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHostContext is an interface to support dynamic dispatch.
type IHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostContext differentiates from other interfaces.
	IsHostContext()
}

type HostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostContext() *HostContext {
	var p = new(HostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_host
	return p
}

func (*HostContext) IsHostContext() {}

func NewHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostContext {
	var p = new(HostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_host

	return p
}

func (s *HostContext) GetParser() antlr.Parser { return s.parser }

func (s *HostContext) Ipv6address() IIpv6addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpv6addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpv6addressContext)
}

func (s *HostContext) Ipv4addressorregname() IIpv4addressorregnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpv4addressorregnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpv4addressorregnameContext)
}

func (s *HostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHost(s)
	}
}

func (s *HostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHost(s)
	}
}

func (p *httpSpecParser) Host() (localctx IHostContext) {
	localctx = NewHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, httpSpecParserRULE_host)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(httpSpecParserT__18)
		}
		{
			p.SetState(186)
			p.Ipv6address()
		}
		{
			p.SetState(187)
			p.Match(httpSpecParserT__19)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Ipv4addressorregname()
		}

	}

	return localctx
}

// IIpv6addressContext is an interface to support dynamic dispatch.
type IIpv6addressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpv6addressContext differentiates from other interfaces.
	IsIpv6addressContext()
}

type Ipv6addressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv6addressContext() *Ipv6addressContext {
	var p = new(Ipv6addressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_ipv6address
	return p
}

func (*Ipv6addressContext) IsIpv6addressContext() {}

func NewIpv6addressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv6addressContext {
	var p = new(Ipv6addressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_ipv6address

	return p
}

func (s *Ipv6addressContext) GetParser() antlr.Parser { return s.parser }
func (s *Ipv6addressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv6addressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv6addressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterIpv6address(s)
	}
}

func (s *Ipv6addressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitIpv6address(s)
	}
}

func (p *httpSpecParser) Ipv6address() (localctx IIpv6addressContext) {
	localctx = NewIpv6addressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, httpSpecParserRULE_ipv6address)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__0)|(1<<httpSpecParserT__1)|(1<<httpSpecParserT__2)|(1<<httpSpecParserT__3)|(1<<httpSpecParserT__4)|(1<<httpSpecParserT__5)|(1<<httpSpecParserT__6)|(1<<httpSpecParserT__7)|(1<<httpSpecParserT__8)|(1<<httpSpecParserT__9)|(1<<httpSpecParserT__10)|(1<<httpSpecParserT__11)|(1<<httpSpecParserT__12)|(1<<httpSpecParserT__13)|(1<<httpSpecParserT__14)|(1<<httpSpecParserT__15)|(1<<httpSpecParserT__16)|(1<<httpSpecParserT__17)|(1<<httpSpecParserT__18)|(1<<httpSpecParserT__21)|(1<<httpSpecParserT__22)|(1<<httpSpecParserT__23)|(1<<httpSpecParserT__24)|(1<<httpSpecParserT__25)|(1<<httpSpecParserT__26)|(1<<httpSpecParserT__27)|(1<<httpSpecParserT__28)|(1<<httpSpecParserT__29)|(1<<httpSpecParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(httpSpecParserT__31-32))|(1<<(httpSpecParserT__32-32))|(1<<(httpSpecParserT__33-32))|(1<<(httpSpecParserT__34-32))|(1<<(httpSpecParserT__35-32))|(1<<(httpSpecParserT__36-32))|(1<<(httpSpecParserNEWLINE-32))|(1<<(httpSpecParserNEWLINEWITHINDENT-32))|(1<<(httpSpecParserLINETAIL-32))|(1<<(httpSpecParserWHITESPACE-32))|(1<<(httpSpecParserOPTIONALWHITESPACE-32))|(1<<(httpSpecParserREQUIREDWHITESPACE-32))|(1<<(httpSpecParserINPUT-32))|(1<<(httpSpecParserALPHA-32))|(1<<(httpSpecParserDIGIT-32))|(1<<(httpSpecParserID-32))|(1<<(httpSpecParserLINECOMMENT-32))|(1<<(httpSpecParserREQUESTSEPARATOR-32)))) != 0) {
		{
			p.SetState(192)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == httpSpecParserT__19 || _la == httpSpecParserT__20 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpv4addressorregnameContext is an interface to support dynamic dispatch.
type IIpv4addressorregnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpv4addressorregnameContext differentiates from other interfaces.
	IsIpv4addressorregnameContext()
}

type Ipv4addressorregnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv4addressorregnameContext() *Ipv4addressorregnameContext {
	var p = new(Ipv4addressorregnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_ipv4addressorregname
	return p
}

func (*Ipv4addressorregnameContext) IsIpv4addressorregnameContext() {}

func NewIpv4addressorregnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv4addressorregnameContext {
	var p = new(Ipv4addressorregnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_ipv4addressorregname

	return p
}

func (s *Ipv4addressorregnameContext) GetParser() antlr.Parser { return s.parser }
func (s *Ipv4addressorregnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv4addressorregnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv4addressorregnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterIpv4addressorregname(s)
	}
}

func (s *Ipv4addressorregnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitIpv4addressorregname(s)
	}
}

func (p *httpSpecParser) Ipv4addressorregname() (localctx IIpv4addressorregnameContext) {
	localctx = NewIpv4addressorregnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, httpSpecParserRULE_ipv4addressorregname)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(197)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__11)|(1<<httpSpecParserT__12)|(1<<httpSpecParserT__17)|(1<<httpSpecParserT__20))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IAbsolutepathContext is an interface to support dynamic dispatch.
type IAbsolutepathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsolutepathContext differentiates from other interfaces.
	IsAbsolutepathContext()
}

type AbsolutepathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsolutepathContext() *AbsolutepathContext {
	var p = new(AbsolutepathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_absolutepath
	return p
}

func (*AbsolutepathContext) IsAbsolutepathContext() {}

func NewAbsolutepathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsolutepathContext {
	var p = new(AbsolutepathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_absolutepath

	return p
}

func (s *AbsolutepathContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsolutepathContext) AllPathseparator() []IPathseparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPathseparatorContext)(nil)).Elem())
	var tst = make([]IPathseparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPathseparatorContext)
		}
	}

	return tst
}

func (s *AbsolutepathContext) Pathseparator(i int) IPathseparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathseparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPathseparatorContext)
}

func (s *AbsolutepathContext) AllSegment() []ISegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISegmentContext)(nil)).Elem())
	var tst = make([]ISegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISegmentContext)
		}
	}

	return tst
}

func (s *AbsolutepathContext) Segment(i int) ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *AbsolutepathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsolutepathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsolutepathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterAbsolutepath(s)
	}
}

func (s *AbsolutepathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitAbsolutepath(s)
	}
}

func (p *httpSpecParser) Absolutepath() (localctx IAbsolutepathContext) {
	localctx = NewAbsolutepathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, httpSpecParserRULE_absolutepath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(httpSpecParserT__20)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == httpSpecParserT__20 || _la == httpSpecParserNEWLINEWITHINDENT {
			{
				p.SetState(203)
				p.Pathseparator()
			}
			{
				p.SetState(204)
				p.Segment()
			}

			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IPathseparatorContext is an interface to support dynamic dispatch.
type IPathseparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathseparatorContext differentiates from other interfaces.
	IsPathseparatorContext()
}

type PathseparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathseparatorContext() *PathseparatorContext {
	var p = new(PathseparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_pathseparator
	return p
}

func (*PathseparatorContext) IsPathseparatorContext() {}

func NewPathseparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathseparatorContext {
	var p = new(PathseparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_pathseparator

	return p
}

func (s *PathseparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PathseparatorContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINEWITHINDENT, 0)
}

func (s *PathseparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathseparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathseparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterPathseparator(s)
	}
}

func (s *PathseparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitPathseparator(s)
	}
}

func (p *httpSpecParser) Pathseparator() (localctx IPathseparatorContext) {
	localctx = NewPathseparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, httpSpecParserRULE_pathseparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		_la = p.GetTokenStream().LA(1)

		if !(_la == httpSpecParserT__20 || _la == httpSpecParserNEWLINEWITHINDENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_segment
	return p
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }
func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *httpSpecParser) Segment() (localctx ISegmentContext) {
	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, httpSpecParserRULE_segment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(214)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__11)|(1<<httpSpecParserT__12)|(1<<httpSpecParserT__20))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINEWITHINDENT, 0)
}

func (s *QueryContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *httpSpecParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, httpSpecParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(220)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == httpSpecParserT__12 {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserNEWLINEWITHINDENT {
		{
			p.SetState(226)
			p.Match(httpSpecParserNEWLINEWITHINDENT)
		}
		{
			p.SetState(227)
			p.Query()
		}

	}

	return localctx
}

// IRequestfragmentContext is an interface to support dynamic dispatch.
type IRequestfragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestfragmentContext differentiates from other interfaces.
	IsRequestfragmentContext()
}

type RequestfragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestfragmentContext() *RequestfragmentContext {
	var p = new(RequestfragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_requestfragment
	return p
}

func (*RequestfragmentContext) IsRequestfragmentContext() {}

func NewRequestfragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestfragmentContext {
	var p = new(RequestfragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_requestfragment

	return p
}

func (s *RequestfragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestfragmentContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINEWITHINDENT, 0)
}

func (s *RequestfragmentContext) Requestfragment() IRequestfragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestfragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestfragmentContext)
}

func (s *RequestfragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestfragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestfragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterRequestfragment(s)
	}
}

func (s *RequestfragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitRequestfragment(s)
	}
}

func (p *httpSpecParser) Requestfragment() (localctx IRequestfragmentContext) {
	localctx = NewRequestfragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, httpSpecParserRULE_requestfragment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == httpSpecParserT__11 {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserNEWLINEWITHINDENT {
		{
			p.SetState(236)
			p.Match(httpSpecParserNEWLINEWITHINDENT)
		}
		{
			p.SetState(237)
			p.Requestfragment()
		}

	}

	return localctx
}

// IHeadersContext is an interface to support dynamic dispatch.
type IHeadersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeadersContext differentiates from other interfaces.
	IsHeadersContext()
}

type HeadersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeadersContext() *HeadersContext {
	var p = new(HeadersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_headers
	return p
}

func (*HeadersContext) IsHeadersContext() {}

func NewHeadersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadersContext {
	var p = new(HeadersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_headers

	return p
}

func (s *HeadersContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadersContext) AllHeaderfield() []IHeaderfieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem())
	var tst = make([]IHeaderfieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderfieldContext)
		}
	}

	return tst
}

func (s *HeadersContext) Headerfield(i int) IHeaderfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderfieldContext)
}

func (s *HeadersContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserNEWLINE)
}

func (s *HeadersContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, i)
}

func (s *HeadersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHeaders(s)
	}
}

func (s *HeadersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHeaders(s)
	}
}

func (p *httpSpecParser) Headers() (localctx IHeadersContext) {
	localctx = NewHeadersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, httpSpecParserRULE_headers)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(240)
				p.Headerfield()
			}
			{
				p.SetState(241)
				p.Match(httpSpecParserNEWLINE)
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// IHeaderfieldContext is an interface to support dynamic dispatch.
type IHeaderfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderfieldContext differentiates from other interfaces.
	IsHeaderfieldContext()
}

type HeaderfieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderfieldContext() *HeaderfieldContext {
	var p = new(HeaderfieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_headerfield
	return p
}

func (*HeaderfieldContext) IsHeaderfieldContext() {}

func NewHeaderfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderfieldContext {
	var p = new(HeaderfieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_headerfield

	return p
}

func (s *HeaderfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderfieldContext) Fieldname() IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *HeaderfieldContext) AllOPTIONALWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserOPTIONALWHITESPACE)
}

func (s *HeaderfieldContext) OPTIONALWHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserOPTIONALWHITESPACE, i)
}

func (s *HeaderfieldContext) Fieldvalue() IFieldvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldvalueContext)
}

func (s *HeaderfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHeaderfield(s)
	}
}

func (s *HeaderfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHeaderfield(s)
	}
}

func (p *httpSpecParser) Headerfield() (localctx IHeaderfieldContext) {
	localctx = NewHeaderfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, httpSpecParserRULE_headerfield)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Fieldname()
	}
	{
		p.SetState(249)
		p.Match(httpSpecParserT__17)
	}
	{
		p.SetState(250)
		p.Match(httpSpecParserOPTIONALWHITESPACE)
	}
	{
		p.SetState(251)
		p.Fieldvalue()
	}
	{
		p.SetState(252)
		p.Match(httpSpecParserOPTIONALWHITESPACE)
	}

	return localctx
}

// IFieldnameContext is an interface to support dynamic dispatch.
type IFieldnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldnameContext differentiates from other interfaces.
	IsFieldnameContext()
}

type FieldnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldnameContext() *FieldnameContext {
	var p = new(FieldnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_fieldname
	return p
}

func (*FieldnameContext) IsFieldnameContext() {}

func NewFieldnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldnameContext {
	var p = new(FieldnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_fieldname

	return p
}

func (s *FieldnameContext) GetParser() antlr.Parser { return s.parser }
func (s *FieldnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterFieldname(s)
	}
}

func (s *FieldnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitFieldname(s)
	}
}

func (p *httpSpecParser) Fieldname() (localctx IFieldnameContext) {
	localctx = NewFieldnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, httpSpecParserRULE_fieldname)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__0)|(1<<httpSpecParserT__1)|(1<<httpSpecParserT__2)|(1<<httpSpecParserT__3)|(1<<httpSpecParserT__4)|(1<<httpSpecParserT__5)|(1<<httpSpecParserT__6)|(1<<httpSpecParserT__7)|(1<<httpSpecParserT__8)|(1<<httpSpecParserT__9)|(1<<httpSpecParserT__10)|(1<<httpSpecParserT__11)|(1<<httpSpecParserT__12)|(1<<httpSpecParserT__13)|(1<<httpSpecParserT__14)|(1<<httpSpecParserT__15)|(1<<httpSpecParserT__16)|(1<<httpSpecParserT__18)|(1<<httpSpecParserT__19)|(1<<httpSpecParserT__20)|(1<<httpSpecParserT__21)|(1<<httpSpecParserT__22)|(1<<httpSpecParserT__23)|(1<<httpSpecParserT__24)|(1<<httpSpecParserT__25)|(1<<httpSpecParserT__26)|(1<<httpSpecParserT__27)|(1<<httpSpecParserT__28)|(1<<httpSpecParserT__29)|(1<<httpSpecParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(httpSpecParserT__31-32))|(1<<(httpSpecParserT__32-32))|(1<<(httpSpecParserT__33-32))|(1<<(httpSpecParserT__34-32))|(1<<(httpSpecParserT__35-32))|(1<<(httpSpecParserT__36-32))|(1<<(httpSpecParserNEWLINE-32))|(1<<(httpSpecParserNEWLINEWITHINDENT-32))|(1<<(httpSpecParserLINETAIL-32))|(1<<(httpSpecParserWHITESPACE-32))|(1<<(httpSpecParserOPTIONALWHITESPACE-32))|(1<<(httpSpecParserREQUIREDWHITESPACE-32))|(1<<(httpSpecParserINPUT-32))|(1<<(httpSpecParserALPHA-32))|(1<<(httpSpecParserDIGIT-32))|(1<<(httpSpecParserID-32))|(1<<(httpSpecParserLINECOMMENT-32))|(1<<(httpSpecParserREQUESTSEPARATOR-32)))) != 0) {
		{
			p.SetState(254)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == httpSpecParserT__17 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldvalueContext is an interface to support dynamic dispatch.
type IFieldvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldvalueContext differentiates from other interfaces.
	IsFieldvalueContext()
}

type FieldvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldvalueContext() *FieldvalueContext {
	var p = new(FieldvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_fieldvalue
	return p
}

func (*FieldvalueContext) IsFieldvalueContext() {}

func NewFieldvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldvalueContext {
	var p = new(FieldvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_fieldvalue

	return p
}

func (s *FieldvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldvalueContext) LINETAIL() antlr.TerminalNode {
	return s.GetToken(httpSpecParserLINETAIL, 0)
}

func (s *FieldvalueContext) NEWLINEWITHINDENT() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINEWITHINDENT, 0)
}

func (s *FieldvalueContext) Fieldvalue() IFieldvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldvalueContext)
}

func (s *FieldvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterFieldvalue(s)
	}
}

func (s *FieldvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitFieldvalue(s)
	}
}

func (p *httpSpecParser) Fieldvalue() (localctx IFieldvalueContext) {
	localctx = NewFieldvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, httpSpecParserRULE_fieldvalue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(httpSpecParserLINETAIL)
	}

	{
		p.SetState(260)
		p.Match(httpSpecParserNEWLINEWITHINDENT)
	}
	{
		p.SetState(261)
		p.Fieldvalue()
	}

	return localctx
}

// IMessagebodyContext is an interface to support dynamic dispatch.
type IMessagebodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagebodyContext differentiates from other interfaces.
	IsMessagebodyContext()
}

type MessagebodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagebodyContext() *MessagebodyContext {
	var p = new(MessagebodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_messagebody
	return p
}

func (*MessagebodyContext) IsMessagebodyContext() {}

func NewMessagebodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagebodyContext {
	var p = new(MessagebodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_messagebody

	return p
}

func (s *MessagebodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagebodyContext) Messages() IMessagesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagesContext)
}

func (s *MessagebodyContext) Multipartformdata() IMultipartformdataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartformdataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartformdataContext)
}

func (s *MessagebodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagebodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagebodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMessagebody(s)
	}
}

func (s *MessagebodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMessagebody(s)
	}
}

func (p *httpSpecParser) Messagebody() (localctx IMessagebodyContext) {
	localctx = NewMessagebodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, httpSpecParserRULE_messagebody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Messages()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Multipartformdata()
		}

	}

	return localctx
}

// IMessagesContext is an interface to support dynamic dispatch.
type IMessagesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagesContext differentiates from other interfaces.
	IsMessagesContext()
}

type MessagesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagesContext() *MessagesContext {
	var p = new(MessagesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_messages
	return p
}

func (*MessagesContext) IsMessagesContext() {}

func NewMessagesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagesContext {
	var p = new(MessagesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_messages

	return p
}

func (s *MessagesContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagesContext) AllMessageline() []IMessagelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessagelineContext)(nil)).Elem())
	var tst = make([]IMessagelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessagelineContext)
		}
	}

	return tst
}

func (s *MessagesContext) Messageline(i int) IMessagelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessagelineContext)
}

func (s *MessagesContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, 0)
}

func (s *MessagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMessages(s)
	}
}

func (s *MessagesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMessages(s)
	}
}

func (p *httpSpecParser) Messages() (localctx IMessagesContext) {
	localctx = NewMessagesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, httpSpecParserRULE_messages)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Messageline()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpSpecParserNEWLINE {
		{
			p.SetState(268)
			p.Match(httpSpecParserNEWLINE)
		}
		{
			p.SetState(269)
			p.Messageline()
		}

	}

	return localctx
}

// IMessagelineContext is an interface to support dynamic dispatch.
type IMessagelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessagelineContext differentiates from other interfaces.
	IsMessagelineContext()
}

type MessagelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessagelineContext() *MessagelineContext {
	var p = new(MessagelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_messageline
	return p
}

func (*MessagelineContext) IsMessagelineContext() {}

func NewMessagelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessagelineContext {
	var p = new(MessagelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_messageline

	return p
}

func (s *MessagelineContext) GetParser() antlr.Parser { return s.parser }

func (s *MessagelineContext) LINETAIL() antlr.TerminalNode {
	return s.GetToken(httpSpecParserLINETAIL, 0)
}

func (s *MessagelineContext) Inputfileref() IInputfilerefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputfilerefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputfilerefContext)
}

func (s *MessagelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessagelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessagelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMessageline(s)
	}
}

func (s *MessagelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMessageline(s)
	}
}

func (p *httpSpecParser) Messageline() (localctx IMessagelineContext) {
	localctx = NewMessagelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, httpSpecParserRULE_messageline)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpSpecParserT__0, httpSpecParserT__1, httpSpecParserT__2, httpSpecParserT__3, httpSpecParserT__4, httpSpecParserT__5, httpSpecParserT__6, httpSpecParserT__7, httpSpecParserT__8, httpSpecParserT__9, httpSpecParserT__10, httpSpecParserT__11, httpSpecParserT__12, httpSpecParserT__13, httpSpecParserT__14, httpSpecParserT__15, httpSpecParserT__16, httpSpecParserT__17, httpSpecParserT__18, httpSpecParserT__19, httpSpecParserT__20, httpSpecParserT__24, httpSpecParserT__25, httpSpecParserT__26, httpSpecParserT__27, httpSpecParserT__28, httpSpecParserT__29, httpSpecParserT__30, httpSpecParserT__31, httpSpecParserT__32, httpSpecParserT__33, httpSpecParserT__34, httpSpecParserT__35, httpSpecParserT__36, httpSpecParserNEWLINE, httpSpecParserNEWLINEWITHINDENT, httpSpecParserLINETAIL, httpSpecParserWHITESPACE, httpSpecParserOPTIONALWHITESPACE, httpSpecParserREQUIREDWHITESPACE, httpSpecParserINPUT, httpSpecParserALPHA, httpSpecParserDIGIT, httpSpecParserID, httpSpecParserLINECOMMENT, httpSpecParserREQUESTSEPARATOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__21)|(1<<httpSpecParserT__22)|(1<<httpSpecParserT__23))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(273)
			p.Match(httpSpecParserLINETAIL)
		}

	case httpSpecParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Inputfileref()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInputfilerefContext is an interface to support dynamic dispatch.
type IInputfilerefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputfilerefContext differentiates from other interfaces.
	IsInputfilerefContext()
}

type InputfilerefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputfilerefContext() *InputfilerefContext {
	var p = new(InputfilerefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_inputfileref
	return p
}

func (*InputfilerefContext) IsInputfilerefContext() {}

func NewInputfilerefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputfilerefContext {
	var p = new(InputfilerefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_inputfileref

	return p
}

func (s *InputfilerefContext) GetParser() antlr.Parser { return s.parser }

func (s *InputfilerefContext) REQUIREDWHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUIREDWHITESPACE, 0)
}

func (s *InputfilerefContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *InputfilerefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputfilerefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputfilerefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterInputfileref(s)
	}
}

func (s *InputfilerefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitInputfileref(s)
	}
}

func (p *httpSpecParser) Inputfileref() (localctx IInputfilerefContext) {
	localctx = NewInputfilerefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, httpSpecParserRULE_inputfileref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(httpSpecParserT__21)
	}
	{
		p.SetState(278)
		p.Match(httpSpecParserREQUIREDWHITESPACE)
	}
	{
		p.SetState(279)
		p.Filepath()
	}

	return localctx
}

// IFilepathContext is an interface to support dynamic dispatch.
type IFilepathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilepathContext differentiates from other interfaces.
	IsFilepathContext()
}

type FilepathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilepathContext() *FilepathContext {
	var p = new(FilepathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_filepath
	return p
}

func (*FilepathContext) IsFilepathContext() {}

func NewFilepathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilepathContext {
	var p = new(FilepathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_filepath

	return p
}

func (s *FilepathContext) GetParser() antlr.Parser { return s.parser }

func (s *FilepathContext) LINETAIL() antlr.TerminalNode {
	return s.GetToken(httpSpecParserLINETAIL, 0)
}

func (s *FilepathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilepathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilepathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterFilepath(s)
	}
}

func (s *FilepathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitFilepath(s)
	}
}

func (p *httpSpecParser) Filepath() (localctx IFilepathContext) {
	localctx = NewFilepathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, httpSpecParserRULE_filepath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(httpSpecParserLINETAIL)
	}

	return localctx
}

// IMultipartformdataContext is an interface to support dynamic dispatch.
type IMultipartformdataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultipartformdataContext differentiates from other interfaces.
	IsMultipartformdataContext()
}

type MultipartformdataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipartformdataContext() *MultipartformdataContext {
	var p = new(MultipartformdataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_multipartformdata
	return p
}

func (*MultipartformdataContext) IsMultipartformdataContext() {}

func NewMultipartformdataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipartformdataContext {
	var p = new(MultipartformdataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_multipartformdata

	return p
}

func (s *MultipartformdataContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipartformdataContext) Multipartfield() IMultipartfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartfieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartfieldContext)
}

func (s *MultipartformdataContext) Boundary() IBoundaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundaryContext)
}

func (s *MultipartformdataContext) Multipartformdata() IMultipartformdataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultipartformdataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultipartformdataContext)
}

func (s *MultipartformdataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipartformdataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipartformdataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMultipartformdata(s)
	}
}

func (s *MultipartformdataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMultipartformdata(s)
	}
}

func (p *httpSpecParser) Multipartformdata() (localctx IMultipartformdataContext) {
	localctx = NewMultipartformdataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, httpSpecParserRULE_multipartformdata)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Multipartfield()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(284)
			p.Multipartformdata()
		}

	}
	{
		p.SetState(287)
		p.Boundary()
	}

	return localctx
}

// IMultipartfieldContext is an interface to support dynamic dispatch.
type IMultipartfieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultipartfieldContext differentiates from other interfaces.
	IsMultipartfieldContext()
}

type MultipartfieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipartfieldContext() *MultipartfieldContext {
	var p = new(MultipartfieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_multipartfield
	return p
}

func (*MultipartfieldContext) IsMultipartfieldContext() {}

func NewMultipartfieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipartfieldContext {
	var p = new(MultipartfieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_multipartfield

	return p
}

func (s *MultipartfieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipartfieldContext) Boundary() IBoundaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundaryContext)
}

func (s *MultipartfieldContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserNEWLINE)
}

func (s *MultipartfieldContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserNEWLINE, i)
}

func (s *MultipartfieldContext) AllHeaderfield() []IHeaderfieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem())
	var tst = make([]IHeaderfieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderfieldContext)
		}
	}

	return tst
}

func (s *MultipartfieldContext) Headerfield(i int) IHeaderfieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderfieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderfieldContext)
}

func (s *MultipartfieldContext) Messages() IMessagesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessagesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessagesContext)
}

func (s *MultipartfieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipartfieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipartfieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterMultipartfield(s)
	}
}

func (s *MultipartfieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitMultipartfield(s)
	}
}

func (p *httpSpecParser) Multipartfield() (localctx IMultipartfieldContext) {
	localctx = NewMultipartfieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, httpSpecParserRULE_multipartfield)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Boundary()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.Headerfield()
			}
			{
				p.SetState(291)
				p.Match(httpSpecParserNEWLINE)
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}
	{
		p.SetState(298)
		p.Match(httpSpecParserNEWLINE)
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(299)
			p.Messages()
		}

	}

	return localctx
}

// IBoundaryContext is an interface to support dynamic dispatch.
type IBoundaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundaryContext differentiates from other interfaces.
	IsBoundaryContext()
}

type BoundaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundaryContext() *BoundaryContext {
	var p = new(BoundaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_boundary
	return p
}

func (*BoundaryContext) IsBoundaryContext() {}

func NewBoundaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundaryContext {
	var p = new(BoundaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_boundary

	return p
}

func (s *BoundaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundaryContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserDIGIT)
}

func (s *BoundaryContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserDIGIT, i)
}

func (s *BoundaryContext) AllALPHA() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserALPHA)
}

func (s *BoundaryContext) ALPHA(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserALPHA, i)
}

func (s *BoundaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterBoundary(s)
	}
}

func (s *BoundaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitBoundary(s)
	}
}

func (p *httpSpecParser) Boundary() (localctx IBoundaryContext) {
	localctx = NewBoundaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, httpSpecParserRULE_boundary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(302)
				_la = p.GetTokenStream().LA(1)

				if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpSpecParserT__10)|(1<<httpSpecParserT__11)|(1<<httpSpecParserT__17)|(1<<httpSpecParserT__20)|(1<<httpSpecParserT__24)|(1<<httpSpecParserT__25)|(1<<httpSpecParserT__26)|(1<<httpSpecParserT__27)|(1<<httpSpecParserT__28)|(1<<httpSpecParserT__29)|(1<<httpSpecParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(httpSpecParserT__31-32))|(1<<(httpSpecParserALPHA-32))|(1<<(httpSpecParserDIGIT-32)))) != 0)) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}

// IResponsehandlerContext is an interface to support dynamic dispatch.
type IResponsehandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResponsehandlerContext differentiates from other interfaces.
	IsResponsehandlerContext()
}

type ResponsehandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponsehandlerContext() *ResponsehandlerContext {
	var p = new(ResponsehandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_responsehandler
	return p
}

func (*ResponsehandlerContext) IsResponsehandlerContext() {}

func NewResponsehandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponsehandlerContext {
	var p = new(ResponsehandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_responsehandler

	return p
}

func (s *ResponsehandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponsehandlerContext) REQUIREDWHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUIREDWHITESPACE, 0)
}

func (s *ResponsehandlerContext) Handlerscript() IHandlerscriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHandlerscriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHandlerscriptContext)
}

func (s *ResponsehandlerContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *ResponsehandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponsehandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponsehandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterResponsehandler(s)
	}
}

func (s *ResponsehandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitResponsehandler(s)
	}
}

func (p *httpSpecParser) Responsehandler() (localctx IResponsehandlerContext) {
	localctx = NewResponsehandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, httpSpecParserRULE_responsehandler)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(httpSpecParserT__32)
	}
	{
		p.SetState(308)
		p.Match(httpSpecParserREQUIREDWHITESPACE)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpSpecParserT__33:
		{
			p.SetState(309)
			p.Match(httpSpecParserT__33)
		}
		{
			p.SetState(310)
			p.Handlerscript()
		}
		{
			p.SetState(311)
			p.Match(httpSpecParserT__34)
		}

	case httpSpecParserLINETAIL:
		{
			p.SetState(313)
			p.Filepath()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHandlerscriptContext is an interface to support dynamic dispatch.
type IHandlerscriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHandlerscriptContext differentiates from other interfaces.
	IsHandlerscriptContext()
}

type HandlerscriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHandlerscriptContext() *HandlerscriptContext {
	var p = new(HandlerscriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_handlerscript
	return p
}

func (*HandlerscriptContext) IsHandlerscriptContext() {}

func NewHandlerscriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HandlerscriptContext {
	var p = new(HandlerscriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_handlerscript

	return p
}

func (s *HandlerscriptContext) GetParser() antlr.Parser { return s.parser }

func (s *HandlerscriptContext) AllINPUT() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserINPUT)
}

func (s *HandlerscriptContext) INPUT(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserINPUT, i)
}

func (s *HandlerscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandlerscriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HandlerscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterHandlerscript(s)
	}
}

func (s *HandlerscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitHandlerscript(s)
	}
}

func (p *httpSpecParser) Handlerscript() (localctx IHandlerscriptContext) {
	localctx = NewHandlerscriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, httpSpecParserRULE_handlerscript)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpSpecParserINPUT {
		{
			p.SetState(316)
			p.Match(httpSpecParserINPUT)
		}

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IResponserefContext is an interface to support dynamic dispatch.
type IResponserefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResponserefContext differentiates from other interfaces.
	IsResponserefContext()
}

type ResponserefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponserefContext() *ResponserefContext {
	var p = new(ResponserefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_responseref
	return p
}

func (*ResponserefContext) IsResponserefContext() {}

func NewResponserefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponserefContext {
	var p = new(ResponserefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_responseref

	return p
}

func (s *ResponserefContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponserefContext) REQUIREDWHITESPACE() antlr.TerminalNode {
	return s.GetToken(httpSpecParserREQUIREDWHITESPACE, 0)
}

func (s *ResponserefContext) Filepath() IFilepathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilepathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilepathContext)
}

func (s *ResponserefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponserefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponserefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterResponseref(s)
	}
}

func (s *ResponserefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitResponseref(s)
	}
}

func (p *httpSpecParser) Responseref() (localctx IResponserefContext) {
	localctx = NewResponserefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, httpSpecParserRULE_responseref)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(httpSpecParserT__22)
	}
	{
		p.SetState(322)
		p.Match(httpSpecParserREQUIREDWHITESPACE)
	}
	{
		p.SetState(323)
		p.Filepath()
	}

	return localctx
}

// IEnvvariableContext is an interface to support dynamic dispatch.
type IEnvvariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnvvariableContext differentiates from other interfaces.
	IsEnvvariableContext()
}

type EnvvariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvvariableContext() *EnvvariableContext {
	var p = new(EnvvariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpSpecParserRULE_envvariable
	return p
}

func (*EnvvariableContext) IsEnvvariableContext() {}

func NewEnvvariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvvariableContext {
	var p = new(EnvvariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpSpecParserRULE_envvariable

	return p
}

func (s *EnvvariableContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvvariableContext) AllOPTIONALWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(httpSpecParserOPTIONALWHITESPACE)
}

func (s *EnvvariableContext) OPTIONALWHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(httpSpecParserOPTIONALWHITESPACE, i)
}

func (s *EnvvariableContext) ID() antlr.TerminalNode {
	return s.GetToken(httpSpecParserID, 0)
}

func (s *EnvvariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvvariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvvariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.EnterEnvvariable(s)
	}
}

func (s *EnvvariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpSpecListener); ok {
		listenerT.ExitEnvvariable(s)
	}
}

func (p *httpSpecParser) Envvariable() (localctx IEnvvariableContext) {
	localctx = NewEnvvariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, httpSpecParserRULE_envvariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(httpSpecParserT__35)
	}
	{
		p.SetState(326)
		p.Match(httpSpecParserOPTIONALWHITESPACE)
	}
	{
		p.SetState(327)
		p.Match(httpSpecParserID)
	}
	{
		p.SetState(328)
		p.Match(httpSpecParserOPTIONALWHITESPACE)
	}
	{
		p.SetState(329)
		p.Match(httpSpecParserT__36)
	}

	return localctx
}
