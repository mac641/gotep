// Code generated from ./src/lib/parser/httpspec.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 378,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 5, 4, 105, 10, 4, 3, 4, 3, 4, 5, 4, 109,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 7, 6, 116, 10, 6, 12, 6, 14, 6, 119,
	11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 137, 10, 12, 13, 12, 14,
	12, 138, 3, 13, 6, 13, 142, 10, 13, 13, 13, 14, 13, 143, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 159, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 224, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 232, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 237, 10,
	26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 6, 28, 245, 10, 28, 13, 28,
	14, 28, 246, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 254, 10, 29, 3,
	30, 6, 30, 257, 10, 30, 13, 30, 14, 30, 258, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 6, 32, 270, 10, 32, 13, 32, 14, 32,
	271, 3, 32, 3, 32, 6, 32, 276, 10, 32, 13, 32, 14, 32, 277, 3, 33, 3, 33,
	3, 33, 3, 33, 5, 33, 284, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 290,
	10, 34, 3, 35, 6, 35, 293, 10, 35, 13, 35, 14, 35, 294, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37,
	309, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 6, 44, 326, 10, 44, 13, 44,
	14, 44, 327, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 334, 10, 44, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 5, 47, 361, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 2, 2, 50, 3, 2, 5, 2, 7, 3, 9, 4, 11, 5, 13, 2, 15, 2, 17, 6, 19, 7,
	21, 8, 23, 9, 25, 10, 27, 11, 29, 12, 31, 13, 33, 2, 35, 2, 37, 14, 39,
	15, 41, 16, 43, 17, 45, 18, 47, 19, 49, 20, 51, 21, 53, 22, 55, 23, 57,
	24, 59, 25, 61, 26, 63, 27, 65, 28, 67, 29, 69, 30, 71, 31, 73, 32, 75,
	2, 77, 2, 79, 2, 81, 33, 83, 34, 85, 35, 87, 36, 89, 37, 91, 38, 93, 39,
	95, 40, 97, 41, 3, 2, 10, 4, 2, 12, 12, 15, 15, 4, 2, 67, 92, 99, 124,
	3, 2, 50, 59, 5, 2, 11, 11, 14, 14, 34, 34, 5, 2, 37, 37, 49, 49, 65, 65,
	3, 2, 37, 37, 3, 2, 65, 65, 3, 2, 60, 60, 2, 399, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 3, 99, 3,
	2, 2, 2, 5, 101, 3, 2, 2, 2, 7, 108, 3, 2, 2, 2, 9, 110, 3, 2, 2, 2, 11,
	117, 3, 2, 2, 2, 13, 122, 3, 2, 2, 2, 15, 124, 3, 2, 2, 2, 17, 126, 3,
	2, 2, 2, 19, 128, 3, 2, 2, 2, 21, 130, 3, 2, 2, 2, 23, 136, 3, 2, 2, 2,
	25, 141, 3, 2, 2, 2, 27, 147, 3, 2, 2, 2, 29, 149, 3, 2, 2, 2, 31, 158,
	3, 2, 2, 2, 33, 162, 3, 2, 2, 2, 35, 164, 3, 2, 2, 2, 37, 166, 3, 2, 2,
	2, 39, 168, 3, 2, 2, 2, 41, 170, 3, 2, 2, 2, 43, 172, 3, 2, 2, 2, 45, 174,
	3, 2, 2, 2, 47, 223, 3, 2, 2, 2, 49, 225, 3, 2, 2, 2, 51, 227, 3, 2, 2,
	2, 53, 238, 3, 2, 2, 2, 55, 240, 3, 2, 2, 2, 57, 248, 3, 2, 2, 2, 59, 256,
	3, 2, 2, 2, 61, 260, 3, 2, 2, 2, 63, 263, 3, 2, 2, 2, 65, 279, 3, 2, 2,
	2, 67, 285, 3, 2, 2, 2, 69, 292, 3, 2, 2, 2, 71, 296, 3, 2, 2, 2, 73, 304,
	3, 2, 2, 2, 75, 310, 3, 2, 2, 2, 77, 312, 3, 2, 2, 2, 79, 314, 3, 2, 2,
	2, 81, 316, 3, 2, 2, 2, 83, 318, 3, 2, 2, 2, 85, 320, 3, 2, 2, 2, 87, 325,
	3, 2, 2, 2, 89, 337, 3, 2, 2, 2, 91, 342, 3, 2, 2, 2, 93, 360, 3, 2, 2,
	2, 95, 362, 3, 2, 2, 2, 97, 368, 3, 2, 2, 2, 99, 100, 7, 15, 2, 2, 100,
	4, 3, 2, 2, 2, 101, 102, 7, 12, 2, 2, 102, 6, 3, 2, 2, 2, 103, 105, 5,
	3, 2, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 106, 109, 5, 5, 3, 2, 107, 109, 5, 3, 2, 2, 108, 104, 3, 2, 2, 2, 108,
	107, 3, 2, 2, 2, 109, 8, 3, 2, 2, 2, 110, 111, 5, 7, 4, 2, 111, 112, 5,
	25, 13, 2, 112, 113, 8, 5, 2, 2, 113, 10, 3, 2, 2, 2, 114, 116, 5, 17,
	9, 2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120,
	121, 5, 7, 4, 2, 121, 12, 3, 2, 2, 2, 122, 123, 7, 47, 2, 2, 123, 14, 3,
	2, 2, 2, 124, 125, 7, 97, 2, 2, 125, 16, 3, 2, 2, 2, 126, 127, 10, 2, 2,
	2, 127, 18, 3, 2, 2, 2, 128, 129, 9, 3, 2, 2, 129, 20, 3, 2, 2, 2, 130,
	131, 9, 4, 2, 2, 131, 22, 3, 2, 2, 2, 132, 137, 5, 19, 10, 2, 133, 137,
	5, 21, 11, 2, 134, 137, 5, 13, 7, 2, 135, 137, 5, 15, 8, 2, 136, 132, 3,
	2, 2, 2, 136, 133, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2,
	2, 137, 138, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	24, 3, 2, 2, 2, 140, 142, 9, 5, 2, 2, 141, 140, 3, 2, 2, 2, 142, 143, 3,
	2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 146, 8, 13, 3, 2, 146, 26, 3, 2, 2, 2, 147, 148, 7, 49, 2, 2, 148,
	28, 3, 2, 2, 2, 149, 150, 7, 37, 2, 2, 150, 30, 3, 2, 2, 2, 151, 152, 5,
	29, 15, 2, 152, 153, 5, 11, 6, 2, 153, 159, 3, 2, 2, 2, 154, 155, 5, 27,
	14, 2, 155, 156, 5, 27, 14, 2, 156, 157, 5, 11, 6, 2, 157, 159, 3, 2, 2,
	2, 158, 151, 3, 2, 2, 2, 158, 154, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	161, 8, 16, 4, 2, 161, 32, 3, 2, 2, 2, 162, 163, 7, 44, 2, 2, 163, 34,
	3, 2, 2, 2, 164, 165, 7, 48, 2, 2, 165, 36, 3, 2, 2, 2, 166, 167, 7, 60,
	2, 2, 167, 38, 3, 2, 2, 2, 168, 169, 7, 65, 2, 2, 169, 40, 3, 2, 2, 2,
	170, 171, 7, 93, 2, 2, 171, 42, 3, 2, 2, 2, 172, 173, 7, 95, 2, 2, 173,
	44, 3, 2, 2, 2, 174, 175, 5, 29, 15, 2, 175, 176, 5, 29, 15, 2, 176, 177,
	5, 29, 15, 2, 177, 178, 5, 11, 6, 2, 178, 46, 3, 2, 2, 2, 179, 180, 7,
	73, 2, 2, 180, 181, 7, 71, 2, 2, 181, 224, 7, 86, 2, 2, 182, 183, 7, 74,
	2, 2, 183, 184, 7, 71, 2, 2, 184, 185, 7, 67, 2, 2, 185, 224, 7, 70, 2,
	2, 186, 187, 7, 82, 2, 2, 187, 188, 7, 81, 2, 2, 188, 189, 7, 85, 2, 2,
	189, 224, 7, 86, 2, 2, 190, 191, 7, 82, 2, 2, 191, 192, 7, 87, 2, 2, 192,
	224, 7, 86, 2, 2, 193, 194, 7, 70, 2, 2, 194, 195, 7, 71, 2, 2, 195, 196,
	7, 78, 2, 2, 196, 197, 7, 71, 2, 2, 197, 198, 7, 86, 2, 2, 198, 224, 7,
	71, 2, 2, 199, 200, 7, 69, 2, 2, 200, 201, 7, 81, 2, 2, 201, 202, 7, 80,
	2, 2, 202, 203, 7, 80, 2, 2, 203, 204, 7, 71, 2, 2, 204, 205, 7, 69, 2,
	2, 205, 224, 7, 86, 2, 2, 206, 207, 7, 82, 2, 2, 207, 208, 7, 67, 2, 2,
	208, 209, 7, 86, 2, 2, 209, 210, 7, 69, 2, 2, 210, 224, 7, 74, 2, 2, 211,
	212, 7, 81, 2, 2, 212, 213, 7, 82, 2, 2, 213, 214, 7, 86, 2, 2, 214, 215,
	7, 75, 2, 2, 215, 216, 7, 81, 2, 2, 216, 217, 7, 80, 2, 2, 217, 224, 7,
	85, 2, 2, 218, 219, 7, 86, 2, 2, 219, 220, 7, 84, 2, 2, 220, 221, 7, 67,
	2, 2, 221, 222, 7, 69, 2, 2, 222, 224, 7, 71, 2, 2, 223, 179, 3, 2, 2,
	2, 223, 182, 3, 2, 2, 2, 223, 186, 3, 2, 2, 2, 223, 190, 3, 2, 2, 2, 223,
	193, 3, 2, 2, 2, 223, 199, 3, 2, 2, 2, 223, 206, 3, 2, 2, 2, 223, 211,
	3, 2, 2, 2, 223, 218, 3, 2, 2, 2, 224, 48, 3, 2, 2, 2, 225, 226, 5, 33,
	17, 2, 226, 50, 3, 2, 2, 2, 227, 231, 5, 55, 28, 2, 228, 229, 5, 39, 20,
	2, 229, 230, 5, 65, 33, 2, 230, 232, 3, 2, 2, 2, 231, 228, 3, 2, 2, 2,
	231, 232, 3, 2, 2, 2, 232, 236, 3, 2, 2, 2, 233, 234, 5, 29, 15, 2, 234,
	235, 5, 67, 34, 2, 235, 237, 3, 2, 2, 2, 236, 233, 3, 2, 2, 2, 236, 237,
	3, 2, 2, 2, 237, 52, 3, 2, 2, 2, 238, 239, 10, 6, 2, 2, 239, 54, 3, 2,
	2, 2, 240, 244, 5, 27, 14, 2, 241, 242, 5, 61, 31, 2, 242, 243, 5, 53,
	27, 2, 243, 245, 3, 2, 2, 2, 244, 241, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2,
	246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 56, 3, 2, 2, 2, 248, 249,
	7, 106, 2, 2, 249, 250, 7, 118, 2, 2, 250, 251, 7, 118, 2, 2, 251, 253,
	7, 114, 2, 2, 252, 254, 7, 117, 2, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3,
	2, 2, 2, 254, 58, 3, 2, 2, 2, 255, 257, 5, 21, 11, 2, 256, 255, 3, 2, 2,
	2, 257, 258, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259,
	60, 3, 2, 2, 2, 260, 261, 5, 27, 14, 2, 261, 262, 5, 9, 5, 2, 262, 62,
	3, 2, 2, 2, 263, 264, 7, 74, 2, 2, 264, 265, 7, 86, 2, 2, 265, 266, 7,
	86, 2, 2, 266, 267, 7, 82, 2, 2, 267, 269, 5, 27, 14, 2, 268, 270, 5, 21,
	11, 2, 269, 268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2,
	271, 272, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 275, 5, 35, 18, 2, 274,
	276, 5, 21, 11, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275,
	3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 64, 3, 2, 2, 2, 279, 283, 10, 7,
	2, 2, 280, 281, 5, 9, 5, 2, 281, 282, 5, 65, 33, 2, 282, 284, 3, 2, 2,
	2, 283, 280, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 66, 3, 2, 2, 2, 285,
	289, 10, 8, 2, 2, 286, 287, 5, 9, 5, 2, 287, 288, 5, 67, 34, 2, 288, 290,
	3, 2, 2, 2, 289, 286, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 68, 3, 2,
	2, 2, 291, 293, 10, 9, 2, 2, 292, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2,
	294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 70, 3, 2, 2, 2, 296, 297,
	5, 69, 35, 2, 297, 298, 5, 37, 19, 2, 298, 299, 5, 25, 13, 2, 299, 300,
	8, 36, 5, 2, 300, 301, 5, 73, 37, 2, 301, 302, 5, 25, 13, 2, 302, 303,
	8, 36, 6, 2, 303, 72, 3, 2, 2, 2, 304, 308, 5, 11, 6, 2, 305, 306, 5, 9,
	5, 2, 306, 307, 5, 73, 37, 2, 307, 309, 3, 2, 2, 2, 308, 305, 3, 2, 2,
	2, 308, 309, 3, 2, 2, 2, 309, 74, 3, 2, 2, 2, 310, 311, 7, 125, 2, 2, 311,
	76, 3, 2, 2, 2, 312, 313, 7, 127, 2, 2, 313, 78, 3, 2, 2, 2, 314, 315,
	7, 39, 2, 2, 315, 80, 3, 2, 2, 2, 316, 317, 7, 62, 2, 2, 317, 82, 3, 2,
	2, 2, 318, 319, 7, 64, 2, 2, 319, 84, 3, 2, 2, 2, 320, 321, 5, 11, 6, 2,
	321, 86, 3, 2, 2, 2, 322, 323, 5, 13, 7, 2, 323, 324, 5, 13, 7, 2, 324,
	326, 3, 2, 2, 2, 325, 322, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 325,
	3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 333, 5, 17,
	9, 2, 330, 331, 5, 13, 7, 2, 331, 332, 5, 13, 7, 2, 332, 334, 3, 2, 2,
	2, 333, 330, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335,
	336, 5, 7, 4, 2, 336, 88, 3, 2, 2, 2, 337, 338, 5, 81, 41, 2, 338, 339,
	5, 25, 13, 2, 339, 340, 8, 45, 7, 2, 340, 341, 5, 85, 43, 2, 341, 90, 3,
	2, 2, 2, 342, 343, 5, 11, 6, 2, 343, 344, 3, 2, 2, 2, 344, 345, 8, 46,
	4, 2, 345, 92, 3, 2, 2, 2, 346, 347, 5, 83, 42, 2, 347, 348, 5, 25, 13,
	2, 348, 349, 8, 47, 8, 2, 349, 350, 5, 75, 38, 2, 350, 351, 5, 79, 40,
	2, 351, 352, 5, 91, 46, 2, 352, 353, 5, 79, 40, 2, 353, 354, 5, 77, 39,
	2, 354, 361, 3, 2, 2, 2, 355, 356, 5, 83, 42, 2, 356, 357, 5, 25, 13, 2,
	357, 358, 8, 47, 9, 2, 358, 359, 5, 85, 43, 2, 359, 361, 3, 2, 2, 2, 360,
	346, 3, 2, 2, 2, 360, 355, 3, 2, 2, 2, 361, 94, 3, 2, 2, 2, 362, 363, 5,
	81, 41, 2, 363, 364, 5, 83, 42, 2, 364, 365, 5, 25, 13, 2, 365, 366, 8,
	48, 10, 2, 366, 367, 5, 85, 43, 2, 367, 96, 3, 2, 2, 2, 368, 369, 5, 75,
	38, 2, 369, 370, 5, 75, 38, 2, 370, 371, 5, 25, 13, 2, 371, 372, 8, 49,
	11, 2, 372, 373, 5, 23, 12, 2, 373, 374, 5, 25, 13, 2, 374, 375, 8, 49,
	12, 2, 375, 376, 5, 77, 39, 2, 376, 377, 5, 77, 39, 2, 377, 98, 3, 2, 2,
	2, 25, 2, 104, 108, 117, 136, 138, 143, 158, 223, 231, 236, 246, 253, 258,
	271, 277, 283, 289, 294, 308, 327, 333, 360, 13, 3, 5, 2, 3, 13, 3, 8,
	2, 2, 3, 36, 4, 3, 36, 5, 3, 45, 6, 3, 47, 7, 3, 47, 8, 3, 48, 9, 3, 49,
	10, 3, 49, 11,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "", "", "'\u002F'", "'\u0023'", "", "'\u003A'",
	"'\u003F'", "'\u005B'", "'\u005D'", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "'\u003C'", "'\u003E'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "INPUT", "ALPHA", "DIGIT",
	"ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT", "COLON", "QUESTIONMARK",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "REQUESTSEPARATOR", "METHOD",
	"ASTERISKFORM", "ORIGINFORM", "SEGMENT", "ABSOLUTEPATH", "SCHEME", "PORT",
	"PATHSEPARATOR", "HTTPVERSION", "QUERY", "REQUESTFRAGMENT", "FIELDNAME",
	"HEADERFIELD", "FIELDVALUE", "LESSTHANSIGN", "GREATERTHANSIGN", "FILEPATH",
	"BOUNDARY", "INPUTFILEREF", "HANDLERSCRIPT", "RESPONSEHANDLER", "RESPONSEREF",
	"ENVVARIABLE",
}

var lexerRuleNames = []string{
	"CR", "LF", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "HYPHEN", "UNDERSCORE",
	"INPUT", "ALPHA", "DIGIT", "ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT",
	"ASTERISK", "FULLSTOP", "COLON", "QUESTIONMARK", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET",
	"REQUESTSEPARATOR", "METHOD", "ASTERISKFORM", "ORIGINFORM", "SEGMENT",
	"ABSOLUTEPATH", "SCHEME", "PORT", "PATHSEPARATOR", "HTTPVERSION", "QUERY",
	"REQUESTFRAGMENT", "FIELDNAME", "HEADERFIELD", "FIELDVALUE", "LEFTCURLYBRACKET",
	"RIGHTCURLYBRACKET", "PERCENTSIGN", "LESSTHANSIGN", "GREATERTHANSIGN",
	"FILEPATH", "BOUNDARY", "INPUTFILEREF", "HANDLERSCRIPT", "RESPONSEHANDLER",
	"RESPONSEREF", "ENVVARIABLE",
}

type httpspecLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewhttpspecLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *httpspecLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpspecLexer(input antlr.CharStream) *httpspecLexer {
	l := new(httpspecLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "httpspec.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// httpspecLexer tokens.
const (
	httpspecLexerNEWLINE            = 1
	httpspecLexerNEWLINEWITHINDENT  = 2
	httpspecLexerLINETAIL           = 3
	httpspecLexerINPUT              = 4
	httpspecLexerALPHA              = 5
	httpspecLexerDIGIT              = 6
	httpspecLexerID                 = 7
	httpspecLexerWHITESPACE         = 8
	httpspecLexerSLASH              = 9
	httpspecLexerHASHTAG            = 10
	httpspecLexerLINECOMMENT        = 11
	httpspecLexerCOLON              = 12
	httpspecLexerQUESTIONMARK       = 13
	httpspecLexerLEFTSQUAREBRACKET  = 14
	httpspecLexerRIGHTSQUAREBRACKET = 15
	httpspecLexerREQUESTSEPARATOR   = 16
	httpspecLexerMETHOD             = 17
	httpspecLexerASTERISKFORM       = 18
	httpspecLexerORIGINFORM         = 19
	httpspecLexerSEGMENT            = 20
	httpspecLexerABSOLUTEPATH       = 21
	httpspecLexerSCHEME             = 22
	httpspecLexerPORT               = 23
	httpspecLexerPATHSEPARATOR      = 24
	httpspecLexerHTTPVERSION        = 25
	httpspecLexerQUERY              = 26
	httpspecLexerREQUESTFRAGMENT    = 27
	httpspecLexerFIELDNAME          = 28
	httpspecLexerHEADERFIELD        = 29
	httpspecLexerFIELDVALUE         = 30
	httpspecLexerLESSTHANSIGN       = 31
	httpspecLexerGREATERTHANSIGN    = 32
	httpspecLexerFILEPATH           = 33
	httpspecLexerBOUNDARY           = 34
	httpspecLexerINPUTFILEREF       = 35
	httpspecLexerHANDLERSCRIPT      = 36
	httpspecLexerRESPONSEHANDLER    = 37
	httpspecLexerRESPONSEREF        = 38
	httpspecLexerENVVARIABLE        = 39
)

var ignoreWs = true

func (l *httpspecLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 3:
		l.NEWLINEWITHINDENT_Action(localctx, actionIndex)

	case 11:
		l.WHITESPACE_Action(localctx, actionIndex)

	case 34:
		l.HEADERFIELD_Action(localctx, actionIndex)

	case 43:
		l.INPUTFILEREF_Action(localctx, actionIndex)

	case 45:
		l.RESPONSEHANDLER_Action(localctx, actionIndex)

	case 46:
		l.RESPONSEREF_Action(localctx, actionIndex)

	case 47:
		l.ENVVARIABLE_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *httpspecLexer) NEWLINEWITHINDENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) WHITESPACE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		if ignoreWs {
			l.Skip()
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) HEADERFIELD_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		ignoreWs = true

	case 3:
		ignoreWs = true

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) INPUTFILEREF_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 4:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) RESPONSEHANDLER_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 5:
		ignoreWs = false

	case 6:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) RESPONSEREF_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 7:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) ENVVARIABLE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 8:
		ignoreWs = true

	case 9:
		ignoreWs = true

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
