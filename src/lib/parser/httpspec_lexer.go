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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 344,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 5, 4,
	99, 10, 4, 3, 4, 3, 4, 5, 4, 103, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	7, 6, 110, 10, 6, 12, 6, 14, 6, 113, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 6, 12, 131, 10, 12, 13, 12, 14, 12, 132, 3, 13, 6, 13, 136, 10, 13,
	13, 13, 14, 13, 137, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 153, 10, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 218,
	10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 226, 10, 26, 3,
	26, 3, 26, 3, 26, 5, 26, 231, 10, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 6, 28, 239, 10, 28, 13, 28, 14, 28, 240, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 5, 29, 248, 10, 29, 3, 30, 6, 30, 251, 10, 30, 13, 30, 14, 30,
	252, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 6,
	32, 264, 10, 32, 13, 32, 14, 32, 265, 3, 32, 3, 32, 6, 32, 270, 10, 32,
	13, 32, 14, 32, 271, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 278, 10, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 5, 34, 284, 10, 34, 3, 35, 6, 35, 287, 10, 35,
	13, 35, 14, 35, 288, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 295, 10, 36, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 43, 6, 43, 312, 10, 43, 13, 43, 14, 43, 313, 3,
	43, 3, 43, 3, 43, 3, 43, 5, 43, 320, 10, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 2, 2, 47, 3,
	2, 5, 2, 7, 3, 9, 4, 11, 5, 13, 2, 15, 2, 17, 6, 19, 7, 21, 8, 23, 9, 25,
	10, 27, 11, 29, 12, 31, 13, 33, 2, 35, 2, 37, 14, 39, 15, 41, 16, 43, 17,
	45, 18, 47, 19, 49, 20, 51, 21, 53, 22, 55, 23, 57, 24, 59, 25, 61, 26,
	63, 27, 65, 28, 67, 29, 69, 30, 71, 31, 73, 32, 75, 33, 77, 34, 79, 35,
	81, 36, 83, 37, 85, 38, 87, 39, 89, 40, 91, 41, 3, 2, 10, 4, 2, 12, 12,
	15, 15, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 11, 14, 14, 34,
	34, 5, 2, 37, 37, 49, 49, 65, 65, 3, 2, 37, 37, 3, 2, 65, 65, 3, 2, 60,
	60, 2, 367, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2,
	75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2,
	2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2,
	2, 2, 91, 3, 2, 2, 2, 3, 93, 3, 2, 2, 2, 5, 95, 3, 2, 2, 2, 7, 102, 3,
	2, 2, 2, 9, 104, 3, 2, 2, 2, 11, 111, 3, 2, 2, 2, 13, 116, 3, 2, 2, 2,
	15, 118, 3, 2, 2, 2, 17, 120, 3, 2, 2, 2, 19, 122, 3, 2, 2, 2, 21, 124,
	3, 2, 2, 2, 23, 130, 3, 2, 2, 2, 25, 135, 3, 2, 2, 2, 27, 141, 3, 2, 2,
	2, 29, 143, 3, 2, 2, 2, 31, 152, 3, 2, 2, 2, 33, 156, 3, 2, 2, 2, 35, 158,
	3, 2, 2, 2, 37, 160, 3, 2, 2, 2, 39, 162, 3, 2, 2, 2, 41, 164, 3, 2, 2,
	2, 43, 166, 3, 2, 2, 2, 45, 168, 3, 2, 2, 2, 47, 217, 3, 2, 2, 2, 49, 219,
	3, 2, 2, 2, 51, 221, 3, 2, 2, 2, 53, 232, 3, 2, 2, 2, 55, 234, 3, 2, 2,
	2, 57, 242, 3, 2, 2, 2, 59, 250, 3, 2, 2, 2, 61, 254, 3, 2, 2, 2, 63, 257,
	3, 2, 2, 2, 65, 273, 3, 2, 2, 2, 67, 279, 3, 2, 2, 2, 69, 286, 3, 2, 2,
	2, 71, 290, 3, 2, 2, 2, 73, 296, 3, 2, 2, 2, 75, 298, 3, 2, 2, 2, 77, 300,
	3, 2, 2, 2, 79, 302, 3, 2, 2, 2, 81, 304, 3, 2, 2, 2, 83, 306, 3, 2, 2,
	2, 85, 311, 3, 2, 2, 2, 87, 323, 3, 2, 2, 2, 89, 328, 3, 2, 2, 2, 91, 334,
	3, 2, 2, 2, 93, 94, 7, 15, 2, 2, 94, 4, 3, 2, 2, 2, 95, 96, 7, 12, 2, 2,
	96, 6, 3, 2, 2, 2, 97, 99, 5, 3, 2, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2,
	2, 2, 99, 100, 3, 2, 2, 2, 100, 103, 5, 5, 3, 2, 101, 103, 5, 3, 2, 2,
	102, 98, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 8, 3, 2, 2, 2, 104, 105,
	5, 7, 4, 2, 105, 106, 5, 25, 13, 2, 106, 107, 8, 5, 2, 2, 107, 10, 3, 2,
	2, 2, 108, 110, 5, 17, 9, 2, 109, 108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2,
	111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113,
	111, 3, 2, 2, 2, 114, 115, 5, 7, 4, 2, 115, 12, 3, 2, 2, 2, 116, 117, 7,
	47, 2, 2, 117, 14, 3, 2, 2, 2, 118, 119, 7, 97, 2, 2, 119, 16, 3, 2, 2,
	2, 120, 121, 10, 2, 2, 2, 121, 18, 3, 2, 2, 2, 122, 123, 9, 3, 2, 2, 123,
	20, 3, 2, 2, 2, 124, 125, 9, 4, 2, 2, 125, 22, 3, 2, 2, 2, 126, 131, 5,
	19, 10, 2, 127, 131, 5, 21, 11, 2, 128, 131, 5, 13, 7, 2, 129, 131, 5,
	15, 8, 2, 130, 126, 3, 2, 2, 2, 130, 127, 3, 2, 2, 2, 130, 128, 3, 2, 2,
	2, 130, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 24, 3, 2, 2, 2, 134, 136, 9, 5, 2, 2, 135, 134, 3,
	2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2,
	2, 138, 139, 3, 2, 2, 2, 139, 140, 8, 13, 3, 2, 140, 26, 3, 2, 2, 2, 141,
	142, 7, 49, 2, 2, 142, 28, 3, 2, 2, 2, 143, 144, 7, 37, 2, 2, 144, 30,
	3, 2, 2, 2, 145, 146, 5, 29, 15, 2, 146, 147, 5, 11, 6, 2, 147, 153, 3,
	2, 2, 2, 148, 149, 5, 27, 14, 2, 149, 150, 5, 27, 14, 2, 150, 151, 5, 11,
	6, 2, 151, 153, 3, 2, 2, 2, 152, 145, 3, 2, 2, 2, 152, 148, 3, 2, 2, 2,
	153, 154, 3, 2, 2, 2, 154, 155, 8, 16, 4, 2, 155, 32, 3, 2, 2, 2, 156,
	157, 7, 44, 2, 2, 157, 34, 3, 2, 2, 2, 158, 159, 7, 48, 2, 2, 159, 36,
	3, 2, 2, 2, 160, 161, 7, 60, 2, 2, 161, 38, 3, 2, 2, 2, 162, 163, 7, 65,
	2, 2, 163, 40, 3, 2, 2, 2, 164, 165, 7, 93, 2, 2, 165, 42, 3, 2, 2, 2,
	166, 167, 7, 95, 2, 2, 167, 44, 3, 2, 2, 2, 168, 169, 5, 29, 15, 2, 169,
	170, 5, 29, 15, 2, 170, 171, 5, 29, 15, 2, 171, 172, 5, 11, 6, 2, 172,
	46, 3, 2, 2, 2, 173, 174, 7, 73, 2, 2, 174, 175, 7, 71, 2, 2, 175, 218,
	7, 86, 2, 2, 176, 177, 7, 74, 2, 2, 177, 178, 7, 71, 2, 2, 178, 179, 7,
	67, 2, 2, 179, 218, 7, 70, 2, 2, 180, 181, 7, 82, 2, 2, 181, 182, 7, 81,
	2, 2, 182, 183, 7, 85, 2, 2, 183, 218, 7, 86, 2, 2, 184, 185, 7, 82, 2,
	2, 185, 186, 7, 87, 2, 2, 186, 218, 7, 86, 2, 2, 187, 188, 7, 70, 2, 2,
	188, 189, 7, 71, 2, 2, 189, 190, 7, 78, 2, 2, 190, 191, 7, 71, 2, 2, 191,
	192, 7, 86, 2, 2, 192, 218, 7, 71, 2, 2, 193, 194, 7, 69, 2, 2, 194, 195,
	7, 81, 2, 2, 195, 196, 7, 80, 2, 2, 196, 197, 7, 80, 2, 2, 197, 198, 7,
	71, 2, 2, 198, 199, 7, 69, 2, 2, 199, 218, 7, 86, 2, 2, 200, 201, 7, 82,
	2, 2, 201, 202, 7, 67, 2, 2, 202, 203, 7, 86, 2, 2, 203, 204, 7, 69, 2,
	2, 204, 218, 7, 74, 2, 2, 205, 206, 7, 81, 2, 2, 206, 207, 7, 82, 2, 2,
	207, 208, 7, 86, 2, 2, 208, 209, 7, 75, 2, 2, 209, 210, 7, 81, 2, 2, 210,
	211, 7, 80, 2, 2, 211, 218, 7, 85, 2, 2, 212, 213, 7, 86, 2, 2, 213, 214,
	7, 84, 2, 2, 214, 215, 7, 67, 2, 2, 215, 216, 7, 69, 2, 2, 216, 218, 7,
	71, 2, 2, 217, 173, 3, 2, 2, 2, 217, 176, 3, 2, 2, 2, 217, 180, 3, 2, 2,
	2, 217, 184, 3, 2, 2, 2, 217, 187, 3, 2, 2, 2, 217, 193, 3, 2, 2, 2, 217,
	200, 3, 2, 2, 2, 217, 205, 3, 2, 2, 2, 217, 212, 3, 2, 2, 2, 218, 48, 3,
	2, 2, 2, 219, 220, 5, 33, 17, 2, 220, 50, 3, 2, 2, 2, 221, 225, 5, 55,
	28, 2, 222, 223, 5, 39, 20, 2, 223, 224, 5, 65, 33, 2, 224, 226, 3, 2,
	2, 2, 225, 222, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 230, 3, 2, 2, 2,
	227, 228, 5, 29, 15, 2, 228, 229, 5, 67, 34, 2, 229, 231, 3, 2, 2, 2, 230,
	227, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 52, 3, 2, 2, 2, 232, 233, 10,
	6, 2, 2, 233, 54, 3, 2, 2, 2, 234, 238, 5, 27, 14, 2, 235, 236, 5, 61,
	31, 2, 236, 237, 5, 53, 27, 2, 237, 239, 3, 2, 2, 2, 238, 235, 3, 2, 2,
	2, 239, 240, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241,
	56, 3, 2, 2, 2, 242, 243, 7, 106, 2, 2, 243, 244, 7, 118, 2, 2, 244, 245,
	7, 118, 2, 2, 245, 247, 7, 114, 2, 2, 246, 248, 7, 117, 2, 2, 247, 246,
	3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 58, 3, 2, 2, 2, 249, 251, 5, 21,
	11, 2, 250, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2,
	252, 253, 3, 2, 2, 2, 253, 60, 3, 2, 2, 2, 254, 255, 5, 27, 14, 2, 255,
	256, 5, 9, 5, 2, 256, 62, 3, 2, 2, 2, 257, 258, 7, 74, 2, 2, 258, 259,
	7, 86, 2, 2, 259, 260, 7, 86, 2, 2, 260, 261, 7, 82, 2, 2, 261, 263, 5,
	27, 14, 2, 262, 264, 5, 21, 11, 2, 263, 262, 3, 2, 2, 2, 264, 265, 3, 2,
	2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2,
	267, 269, 5, 35, 18, 2, 268, 270, 5, 21, 11, 2, 269, 268, 3, 2, 2, 2, 270,
	271, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 64, 3,
	2, 2, 2, 273, 277, 10, 7, 2, 2, 274, 275, 5, 9, 5, 2, 275, 276, 5, 65,
	33, 2, 276, 278, 3, 2, 2, 2, 277, 274, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 66, 3, 2, 2, 2, 279, 283, 10, 8, 2, 2, 280, 281, 5, 9, 5, 2, 281,
	282, 5, 67, 34, 2, 282, 284, 3, 2, 2, 2, 283, 280, 3, 2, 2, 2, 283, 284,
	3, 2, 2, 2, 284, 68, 3, 2, 2, 2, 285, 287, 10, 9, 2, 2, 286, 285, 3, 2,
	2, 2, 287, 288, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2,
	289, 70, 3, 2, 2, 2, 290, 294, 5, 11, 6, 2, 291, 292, 5, 9, 5, 2, 292,
	293, 5, 71, 36, 2, 293, 295, 3, 2, 2, 2, 294, 291, 3, 2, 2, 2, 294, 295,
	3, 2, 2, 2, 295, 72, 3, 2, 2, 2, 296, 297, 7, 125, 2, 2, 297, 74, 3, 2,
	2, 2, 298, 299, 7, 127, 2, 2, 299, 76, 3, 2, 2, 2, 300, 301, 7, 39, 2,
	2, 301, 78, 3, 2, 2, 2, 302, 303, 7, 62, 2, 2, 303, 80, 3, 2, 2, 2, 304,
	305, 7, 64, 2, 2, 305, 82, 3, 2, 2, 2, 306, 307, 5, 11, 6, 2, 307, 84,
	3, 2, 2, 2, 308, 309, 5, 13, 7, 2, 309, 310, 5, 13, 7, 2, 310, 312, 3,
	2, 2, 2, 311, 308, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311, 3, 2, 2,
	2, 313, 314, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 319, 5, 17, 9, 2, 316,
	317, 5, 13, 7, 2, 317, 318, 5, 13, 7, 2, 318, 320, 3, 2, 2, 2, 319, 316,
	3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 322, 5, 7,
	4, 2, 322, 86, 3, 2, 2, 2, 323, 324, 5, 79, 40, 2, 324, 325, 5, 25, 13,
	2, 325, 326, 8, 44, 5, 2, 326, 327, 5, 83, 42, 2, 327, 88, 3, 2, 2, 2,
	328, 329, 5, 79, 40, 2, 329, 330, 5, 81, 41, 2, 330, 331, 5, 25, 13, 2,
	331, 332, 8, 45, 6, 2, 332, 333, 5, 83, 42, 2, 333, 90, 3, 2, 2, 2, 334,
	335, 5, 73, 37, 2, 335, 336, 5, 73, 37, 2, 336, 337, 5, 25, 13, 2, 337,
	338, 8, 46, 7, 2, 338, 339, 5, 23, 12, 2, 339, 340, 5, 25, 13, 2, 340,
	341, 8, 46, 8, 2, 341, 342, 5, 75, 38, 2, 342, 343, 5, 75, 38, 2, 343,
	92, 3, 2, 2, 2, 24, 2, 98, 102, 111, 130, 132, 137, 152, 217, 225, 230,
	240, 247, 252, 265, 271, 277, 283, 288, 294, 313, 319, 9, 3, 5, 2, 3, 13,
	3, 8, 2, 2, 3, 44, 4, 3, 45, 5, 3, 46, 6, 3, 46, 7,
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
	"", "", "", "", "", "'\u007B'", "'\u007D'", "'\u0025'", "'\u003C'", "'\u003E'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "INPUT", "ALPHA", "DIGIT",
	"ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT", "COLON", "QUESTIONMARK",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "REQUESTSEPARATOR", "METHOD",
	"ASTERISKFORM", "ORIGINFORM", "SEGMENT", "ABSOLUTEPATH", "SCHEME", "PORT",
	"PATHSEPARATOR", "HTTPVERSION", "QUERY", "REQUESTFRAGMENT", "FIELDNAME",
	"FIELDVALUE", "LEFTCURLYBRACKET", "RIGHTCURLYBRACKET", "PERCENTSIGN", "LESSTHANSIGN",
	"GREATERTHANSIGN", "FILEPATH", "BOUNDARY", "INPUTFILEREF", "RESPONSEREF",
	"ENVVARIABLE",
}

var lexerRuleNames = []string{
	"CR", "LF", "NEWLINE", "NEWLINEWITHINDENT", "LINETAIL", "HYPHEN", "UNDERSCORE",
	"INPUT", "ALPHA", "DIGIT", "ID", "WHITESPACE", "SLASH", "HASHTAG", "LINECOMMENT",
	"ASTERISK", "FULLSTOP", "COLON", "QUESTIONMARK", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET",
	"REQUESTSEPARATOR", "METHOD", "ASTERISKFORM", "ORIGINFORM", "SEGMENT",
	"ABSOLUTEPATH", "SCHEME", "PORT", "PATHSEPARATOR", "HTTPVERSION", "QUERY",
	"REQUESTFRAGMENT", "FIELDNAME", "FIELDVALUE", "LEFTCURLYBRACKET", "RIGHTCURLYBRACKET",
	"PERCENTSIGN", "LESSTHANSIGN", "GREATERTHANSIGN", "FILEPATH", "BOUNDARY",
	"INPUTFILEREF", "RESPONSEREF", "ENVVARIABLE",
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
	httpspecLexerFIELDVALUE         = 29
	httpspecLexerLEFTCURLYBRACKET   = 30
	httpspecLexerRIGHTCURLYBRACKET  = 31
	httpspecLexerPERCENTSIGN        = 32
	httpspecLexerLESSTHANSIGN       = 33
	httpspecLexerGREATERTHANSIGN    = 34
	httpspecLexerFILEPATH           = 35
	httpspecLexerBOUNDARY           = 36
	httpspecLexerINPUTFILEREF       = 37
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

	case 42:
		l.INPUTFILEREF_Action(localctx, actionIndex)

	case 43:
		l.RESPONSEREF_Action(localctx, actionIndex)

	case 44:
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
func (l *httpspecLexer) INPUTFILEREF_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) RESPONSEREF_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 3:
		ignoreWs = false

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *httpspecLexer) ENVVARIABLE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 4:
		ignoreWs = true

	case 5:
		ignoreWs = true

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
