// Code generated from CAQLLexer.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 80, 739,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 4, 106,
	9, 106, 4, 107, 9, 107, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 294, 10, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 5, 48, 414, 10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 5, 50, 426, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3,
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3,
	56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61,
	3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3,
	63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64,
	3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3,
	67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68,
	3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 7, 70, 555, 10, 70, 12,
	70, 14, 70, 558, 11, 70, 3, 71, 3, 71, 7, 71, 562, 10, 71, 12, 71, 14,
	71, 565, 11, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 6, 71, 572, 10, 71,
	13, 71, 14, 71, 573, 3, 71, 3, 71, 3, 71, 3, 71, 6, 71, 580, 10, 71, 13,
	71, 14, 71, 581, 5, 71, 584, 10, 71, 3, 72, 3, 72, 7, 72, 588, 10, 72,
	12, 72, 14, 72, 591, 11, 72, 3, 72, 5, 72, 594, 10, 72, 3, 72, 3, 72, 6,
	72, 598, 10, 72, 13, 72, 14, 72, 599, 3, 72, 3, 72, 5, 72, 604, 10, 72,
	3, 72, 6, 72, 607, 10, 72, 13, 72, 14, 72, 608, 5, 72, 611, 10, 72, 3,
	73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 7, 74, 622,
	10, 74, 12, 74, 14, 74, 625, 11, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74,
	3, 74, 3, 74, 7, 74, 634, 10, 74, 12, 74, 14, 74, 637, 11, 74, 3, 74, 5,
	74, 640, 10, 74, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75, 646, 10, 75, 12, 75,
	14, 75, 649, 11, 75, 3, 75, 5, 75, 652, 10, 75, 3, 75, 3, 75, 5, 75, 656,
	10, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 7, 76, 664, 10, 76, 12,
	76, 14, 76, 667, 11, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77,
	3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3,
	82, 3, 82, 3, 83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87,
	3, 87, 3, 88, 3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3,
	92, 3, 93, 3, 93, 3, 94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97,
	3, 98, 3, 98, 3, 99, 3, 99, 3, 100, 3, 100, 3, 101, 3, 101, 3, 102, 3,
	102, 3, 103, 3, 103, 3, 104, 3, 104, 3, 105, 3, 105, 3, 106, 3, 106, 3,
	107, 3, 107, 3, 107, 3, 107, 3, 665, 2, 108, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43,
	85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52,
	103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60,
	119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68,
	135, 69, 137, 70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76,
	151, 77, 153, 78, 155, 79, 157, 2, 159, 2, 161, 2, 163, 2, 165, 2, 167,
	2, 169, 2, 171, 2, 173, 2, 175, 2, 177, 2, 179, 2, 181, 2, 183, 2, 185,
	2, 187, 2, 189, 2, 191, 2, 193, 2, 195, 2, 197, 2, 199, 2, 201, 2, 203,
	2, 205, 2, 207, 2, 209, 2, 211, 2, 213, 80, 3, 2, 39, 5, 2, 67, 92, 97,
	97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 51, 59, 3, 2,
	50, 51, 4, 2, 45, 45, 47, 47, 4, 2, 41, 41, 94, 94, 4, 2, 36, 36, 94, 94,
	4, 2, 12, 12, 15, 15, 5, 2, 11, 13, 15, 15, 34, 34, 5, 2, 50, 59, 67, 72,
	99, 104, 3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4,
	2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4,
	2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4,
	2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4,
	2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4,
	2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4,
	2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4,
	2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4,
	2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 2,
	738, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3,
	2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71,
	3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2,
	79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2,
	2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2,
	2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3,
	2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2,
	109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2,
	2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123,
	3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2,
	2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3,
	2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2,
	145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2,
	2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 213, 3, 2, 2, 2, 3, 215,
	3, 2, 2, 2, 5, 217, 3, 2, 2, 2, 7, 220, 3, 2, 2, 2, 9, 223, 3, 2, 2, 2,
	11, 226, 3, 2, 2, 2, 13, 229, 3, 2, 2, 2, 15, 231, 3, 2, 2, 2, 17, 233,
	3, 2, 2, 2, 19, 236, 3, 2, 2, 2, 21, 239, 3, 2, 2, 2, 23, 241, 3, 2, 2,
	2, 25, 243, 3, 2, 2, 2, 27, 245, 3, 2, 2, 2, 29, 247, 3, 2, 2, 2, 31, 249,
	3, 2, 2, 2, 33, 251, 3, 2, 2, 2, 35, 253, 3, 2, 2, 2, 37, 256, 3, 2, 2,
	2, 39, 259, 3, 2, 2, 2, 41, 261, 3, 2, 2, 2, 43, 263, 3, 2, 2, 2, 45, 265,
	3, 2, 2, 2, 47, 267, 3, 2, 2, 2, 49, 269, 3, 2, 2, 2, 51, 271, 3, 2, 2,
	2, 53, 273, 3, 2, 2, 2, 55, 283, 3, 2, 2, 2, 57, 293, 3, 2, 2, 2, 59, 295,
	3, 2, 2, 2, 61, 299, 3, 2, 2, 2, 63, 303, 3, 2, 2, 2, 65, 311, 3, 2, 2,
	2, 67, 316, 3, 2, 2, 2, 69, 325, 3, 2, 2, 2, 71, 331, 3, 2, 2, 2, 73, 338,
	3, 2, 2, 2, 75, 342, 3, 2, 2, 2, 77, 348, 3, 2, 2, 2, 79, 351, 3, 2, 2,
	2, 81, 359, 3, 2, 2, 2, 83, 366, 3, 2, 2, 2, 85, 371, 3, 2, 2, 2, 87, 388,
	3, 2, 2, 2, 89, 392, 3, 2, 2, 2, 91, 397, 3, 2, 2, 2, 93, 403, 3, 2, 2,
	2, 95, 413, 3, 2, 2, 2, 97, 415, 3, 2, 2, 2, 99, 425, 3, 2, 2, 2, 101,
	427, 3, 2, 2, 2, 103, 436, 3, 2, 2, 2, 105, 443, 3, 2, 2, 2, 107, 451,
	3, 2, 2, 2, 109, 458, 3, 2, 2, 2, 111, 472, 3, 2, 2, 2, 113, 477, 3, 2,
	2, 2, 115, 482, 3, 2, 2, 2, 117, 489, 3, 2, 2, 2, 119, 496, 3, 2, 2, 2,
	121, 501, 3, 2, 2, 2, 123, 506, 3, 2, 2, 2, 125, 512, 3, 2, 2, 2, 127,
	520, 3, 2, 2, 2, 129, 526, 3, 2, 2, 2, 131, 533, 3, 2, 2, 2, 133, 536,
	3, 2, 2, 2, 135, 544, 3, 2, 2, 2, 137, 548, 3, 2, 2, 2, 139, 552, 3, 2,
	2, 2, 141, 583, 3, 2, 2, 2, 143, 593, 3, 2, 2, 2, 145, 612, 3, 2, 2, 2,
	147, 639, 3, 2, 2, 2, 149, 641, 3, 2, 2, 2, 151, 659, 3, 2, 2, 2, 153,
	673, 3, 2, 2, 2, 155, 677, 3, 2, 2, 2, 157, 679, 3, 2, 2, 2, 159, 681,
	3, 2, 2, 2, 161, 683, 3, 2, 2, 2, 163, 685, 3, 2, 2, 2, 165, 687, 3, 2,
	2, 2, 167, 689, 3, 2, 2, 2, 169, 691, 3, 2, 2, 2, 171, 693, 3, 2, 2, 2,
	173, 695, 3, 2, 2, 2, 175, 697, 3, 2, 2, 2, 177, 699, 3, 2, 2, 2, 179,
	701, 3, 2, 2, 2, 181, 703, 3, 2, 2, 2, 183, 705, 3, 2, 2, 2, 185, 707,
	3, 2, 2, 2, 187, 709, 3, 2, 2, 2, 189, 711, 3, 2, 2, 2, 191, 713, 3, 2,
	2, 2, 193, 715, 3, 2, 2, 2, 195, 717, 3, 2, 2, 2, 197, 719, 3, 2, 2, 2,
	199, 721, 3, 2, 2, 2, 201, 723, 3, 2, 2, 2, 203, 725, 3, 2, 2, 2, 205,
	727, 3, 2, 2, 2, 207, 729, 3, 2, 2, 2, 209, 731, 3, 2, 2, 2, 211, 733,
	3, 2, 2, 2, 213, 735, 3, 2, 2, 2, 215, 216, 7, 48, 2, 2, 216, 4, 3, 2,
	2, 2, 217, 218, 7, 63, 2, 2, 218, 219, 7, 128, 2, 2, 219, 6, 3, 2, 2, 2,
	220, 221, 7, 35, 2, 2, 221, 222, 7, 128, 2, 2, 222, 8, 3, 2, 2, 2, 223,
	224, 7, 63, 2, 2, 224, 225, 7, 63, 2, 2, 225, 10, 3, 2, 2, 2, 226, 227,
	7, 35, 2, 2, 227, 228, 7, 63, 2, 2, 228, 12, 3, 2, 2, 2, 229, 230, 7, 62,
	2, 2, 230, 14, 3, 2, 2, 2, 231, 232, 7, 64, 2, 2, 232, 16, 3, 2, 2, 2,
	233, 234, 7, 62, 2, 2, 234, 235, 7, 63, 2, 2, 235, 18, 3, 2, 2, 2, 236,
	237, 7, 64, 2, 2, 237, 238, 7, 63, 2, 2, 238, 20, 3, 2, 2, 2, 239, 240,
	7, 45, 2, 2, 240, 22, 3, 2, 2, 2, 241, 242, 7, 47, 2, 2, 242, 24, 3, 2,
	2, 2, 243, 244, 7, 44, 2, 2, 244, 26, 3, 2, 2, 2, 245, 246, 7, 49, 2, 2,
	246, 28, 3, 2, 2, 2, 247, 248, 7, 39, 2, 2, 248, 30, 3, 2, 2, 2, 249, 250,
	7, 65, 2, 2, 250, 32, 3, 2, 2, 2, 251, 252, 7, 60, 2, 2, 252, 34, 3, 2,
	2, 2, 253, 254, 7, 60, 2, 2, 254, 255, 7, 60, 2, 2, 255, 36, 3, 2, 2, 2,
	256, 257, 7, 48, 2, 2, 257, 258, 7, 48, 2, 2, 258, 38, 3, 2, 2, 2, 259,
	260, 7, 46, 2, 2, 260, 40, 3, 2, 2, 2, 261, 262, 7, 42, 2, 2, 262, 42,
	3, 2, 2, 2, 263, 264, 7, 43, 2, 2, 264, 44, 3, 2, 2, 2, 265, 266, 7, 125,
	2, 2, 266, 46, 3, 2, 2, 2, 267, 268, 7, 127, 2, 2, 268, 48, 3, 2, 2, 2,
	269, 270, 7, 93, 2, 2, 270, 50, 3, 2, 2, 2, 271, 272, 7, 95, 2, 2, 272,
	52, 3, 2, 2, 2, 273, 274, 5, 161, 81, 2, 274, 275, 5, 173, 87, 2, 275,
	276, 5, 173, 87, 2, 276, 277, 5, 195, 98, 2, 277, 278, 5, 169, 85, 2, 278,
	279, 5, 173, 87, 2, 279, 280, 5, 161, 81, 2, 280, 281, 5, 199, 100, 2,
	281, 282, 5, 169, 85, 2, 282, 54, 3, 2, 2, 2, 283, 284, 5, 161, 81, 2,
	284, 285, 5, 183, 92, 2, 285, 286, 5, 183, 92, 2, 286, 56, 3, 2, 2, 2,
	287, 288, 5, 161, 81, 2, 288, 289, 5, 187, 94, 2, 289, 290, 5, 167, 84,
	2, 290, 294, 3, 2, 2, 2, 291, 292, 7, 40, 2, 2, 292, 294, 7, 40, 2, 2,
	293, 287, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 58, 3, 2, 2, 2, 295, 296,
	5, 161, 81, 2, 296, 297, 5, 187, 94, 2, 297, 298, 5, 209, 105, 2, 298,
	60, 3, 2, 2, 2, 299, 300, 5, 161, 81, 2, 300, 301, 5, 197, 99, 2, 301,
	302, 5, 165, 83, 2, 302, 62, 3, 2, 2, 2, 303, 304, 5, 165, 83, 2, 304,
	305, 5, 189, 95, 2, 305, 306, 5, 183, 92, 2, 306, 307, 5, 183, 92, 2, 307,
	308, 5, 169, 85, 2, 308, 309, 5, 165, 83, 2, 309, 310, 5, 199, 100, 2,
	310, 64, 3, 2, 2, 2, 311, 312, 5, 167, 84, 2, 312, 313, 5, 169, 85, 2,
	313, 314, 5, 197, 99, 2, 314, 315, 5, 165, 83, 2, 315, 66, 3, 2, 2, 2,
	316, 317, 5, 167, 84, 2, 317, 318, 5, 177, 89, 2, 318, 319, 5, 197, 99,
	2, 319, 320, 5, 199, 100, 2, 320, 321, 5, 177, 89, 2, 321, 322, 5, 187,
	94, 2, 322, 323, 5, 165, 83, 2, 323, 324, 5, 199, 100, 2, 324, 68, 3, 2,
	2, 2, 325, 326, 5, 171, 86, 2, 326, 327, 5, 161, 81, 2, 327, 328, 5, 183,
	92, 2, 328, 329, 5, 197, 99, 2, 329, 330, 5, 169, 85, 2, 330, 70, 3, 2,
	2, 2, 331, 332, 5, 171, 86, 2, 332, 333, 5, 177, 89, 2, 333, 334, 5, 183,
	92, 2, 334, 335, 5, 199, 100, 2, 335, 336, 5, 169, 85, 2, 336, 337, 5,
	195, 98, 2, 337, 72, 3, 2, 2, 2, 338, 339, 5, 171, 86, 2, 339, 340, 5,
	189, 95, 2, 340, 341, 5, 195, 98, 2, 341, 74, 3, 2, 2, 2, 342, 343, 5,
	173, 87, 2, 343, 344, 5, 195, 98, 2, 344, 345, 5, 161, 81, 2, 345, 346,
	5, 191, 96, 2, 346, 347, 5, 175, 88, 2, 347, 76, 3, 2, 2, 2, 348, 349,
	5, 177, 89, 2, 349, 350, 5, 187, 94, 2, 350, 78, 3, 2, 2, 2, 351, 352,
	5, 177, 89, 2, 352, 353, 5, 187, 94, 2, 353, 354, 5, 163, 82, 2, 354, 355,
	5, 189, 95, 2, 355, 356, 5, 201, 101, 2, 356, 357, 5, 187, 94, 2, 357,
	358, 5, 167, 84, 2, 358, 80, 3, 2, 2, 2, 359, 360, 5, 177, 89, 2, 360,
	361, 5, 187, 94, 2, 361, 362, 5, 197, 99, 2, 362, 363, 5, 169, 85, 2, 363,
	364, 5, 195, 98, 2, 364, 365, 5, 199, 100, 2, 365, 82, 3, 2, 2, 2, 366,
	367, 5, 177, 89, 2, 367, 368, 5, 187, 94, 2, 368, 369, 5, 199, 100, 2,
	369, 370, 5, 189, 95, 2, 370, 84, 3, 2, 2, 2, 371, 372, 5, 181, 91, 2,
	372, 373, 7, 97, 2, 2, 373, 374, 5, 197, 99, 2, 374, 375, 5, 175, 88, 2,
	375, 376, 5, 189, 95, 2, 376, 377, 5, 195, 98, 2, 377, 378, 5, 199, 100,
	2, 378, 379, 5, 169, 85, 2, 379, 380, 5, 197, 99, 2, 380, 381, 5, 199,
	100, 2, 381, 382, 7, 97, 2, 2, 382, 383, 5, 191, 96, 2, 383, 384, 5, 161,
	81, 2, 384, 385, 5, 199, 100, 2, 385, 386, 5, 175, 88, 2, 386, 387, 5,
	197, 99, 2, 387, 86, 3, 2, 2, 2, 388, 389, 5, 183, 92, 2, 389, 390, 5,
	169, 85, 2, 390, 391, 5, 199, 100, 2, 391, 88, 3, 2, 2, 2, 392, 393, 5,
	183, 92, 2, 393, 394, 5, 177, 89, 2, 394, 395, 5, 181, 91, 2, 395, 396,
	5, 169, 85, 2, 396, 90, 3, 2, 2, 2, 397, 398, 5, 183, 92, 2, 398, 399,
	5, 177, 89, 2, 399, 400, 5, 185, 93, 2, 400, 401, 5, 177, 89, 2, 401, 402,
	5, 199, 100, 2, 402, 92, 3, 2, 2, 2, 403, 404, 5, 187, 94, 2, 404, 405,
	5, 189, 95, 2, 405, 406, 5, 187, 94, 2, 406, 407, 5, 169, 85, 2, 407, 94,
	3, 2, 2, 2, 408, 409, 5, 187, 94, 2, 409, 410, 5, 189, 95, 2, 410, 411,
	5, 199, 100, 2, 411, 414, 3, 2, 2, 2, 412, 414, 7, 35, 2, 2, 413, 408,
	3, 2, 2, 2, 413, 412, 3, 2, 2, 2, 414, 96, 3, 2, 2, 2, 415, 416, 5, 187,
	94, 2, 416, 417, 5, 201, 101, 2, 417, 418, 5, 183, 92, 2, 418, 419, 5,
	183, 92, 2, 419, 98, 3, 2, 2, 2, 420, 421, 5, 189, 95, 2, 421, 422, 5,
	195, 98, 2, 422, 426, 3, 2, 2, 2, 423, 424, 7, 126, 2, 2, 424, 426, 7,
	126, 2, 2, 425, 420, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 100, 3, 2,
	2, 2, 427, 428, 5, 189, 95, 2, 428, 429, 5, 201, 101, 2, 429, 430, 5, 199,
	100, 2, 430, 431, 5, 163, 82, 2, 431, 432, 5, 189, 95, 2, 432, 433, 5,
	201, 101, 2, 433, 434, 5, 187, 94, 2, 434, 435, 5, 167, 84, 2, 435, 102,
	3, 2, 2, 2, 436, 437, 5, 195, 98, 2, 437, 438, 5, 169, 85, 2, 438, 439,
	5, 185, 93, 2, 439, 440, 5, 189, 95, 2, 440, 441, 5, 203, 102, 2, 441,
	442, 5, 169, 85, 2, 442, 104, 3, 2, 2, 2, 443, 444, 5, 195, 98, 2, 444,
	445, 5, 169, 85, 2, 445, 446, 5, 191, 96, 2, 446, 447, 5, 183, 92, 2, 447,
	448, 5, 161, 81, 2, 448, 449, 5, 165, 83, 2, 449, 450, 5, 169, 85, 2, 450,
	106, 3, 2, 2, 2, 451, 452, 5, 195, 98, 2, 452, 453, 5, 169, 85, 2, 453,
	454, 5, 199, 100, 2, 454, 455, 5, 201, 101, 2, 455, 456, 5, 195, 98, 2,
	456, 457, 5, 187, 94, 2, 457, 108, 3, 2, 2, 2, 458, 459, 5, 197, 99, 2,
	459, 460, 5, 175, 88, 2, 460, 461, 5, 189, 95, 2, 461, 462, 5, 195, 98,
	2, 462, 463, 5, 199, 100, 2, 463, 464, 5, 169, 85, 2, 464, 465, 5, 197,
	99, 2, 465, 466, 5, 199, 100, 2, 466, 467, 7, 97, 2, 2, 467, 468, 5, 191,
	96, 2, 468, 469, 5, 161, 81, 2, 469, 470, 5, 199, 100, 2, 470, 471, 5,
	175, 88, 2, 471, 110, 3, 2, 2, 2, 472, 473, 5, 197, 99, 2, 473, 474, 5,
	189, 95, 2, 474, 475, 5, 195, 98, 2, 475, 476, 5, 199, 100, 2, 476, 112,
	3, 2, 2, 2, 477, 478, 5, 199, 100, 2, 478, 479, 5, 195, 98, 2, 479, 480,
	5, 201, 101, 2, 480, 481, 5, 169, 85, 2, 481, 114, 3, 2, 2, 2, 482, 483,
	5, 201, 101, 2, 483, 484, 5, 191, 96, 2, 484, 485, 5, 167, 84, 2, 485,
	486, 5, 161, 81, 2, 486, 487, 5, 199, 100, 2, 487, 488, 5, 169, 85, 2,
	488, 116, 3, 2, 2, 2, 489, 490, 5, 201, 101, 2, 490, 491, 5, 191, 96, 2,
	491, 492, 5, 197, 99, 2, 492, 493, 5, 169, 85, 2, 493, 494, 5, 195, 98,
	2, 494, 495, 5, 199, 100, 2, 495, 118, 3, 2, 2, 2, 496, 497, 5, 205, 103,
	2, 497, 498, 5, 177, 89, 2, 498, 499, 5, 199, 100, 2, 499, 500, 5, 175,
	88, 2, 500, 120, 3, 2, 2, 2, 501, 502, 5, 181, 91, 2, 502, 503, 5, 169,
	85, 2, 503, 504, 5, 169, 85, 2, 504, 505, 5, 191, 96, 2, 505, 122, 3, 2,
	2, 2, 506, 507, 5, 165, 83, 2, 507, 508, 5, 189, 95, 2, 508, 509, 5, 201,
	101, 2, 509, 510, 5, 187, 94, 2, 510, 511, 5, 199, 100, 2, 511, 124, 3,
	2, 2, 2, 512, 513, 5, 189, 95, 2, 513, 514, 5, 191, 96, 2, 514, 515, 5,
	199, 100, 2, 515, 516, 5, 177, 89, 2, 516, 517, 5, 189, 95, 2, 517, 518,
	5, 187, 94, 2, 518, 519, 5, 197, 99, 2, 519, 126, 3, 2, 2, 2, 520, 521,
	5, 191, 96, 2, 521, 522, 5, 195, 98, 2, 522, 523, 5, 201, 101, 2, 523,
	524, 5, 187, 94, 2, 524, 525, 5, 169, 85, 2, 525, 128, 3, 2, 2, 2, 526,
	527, 5, 197, 99, 2, 527, 528, 5, 169, 85, 2, 528, 529, 5, 161, 81, 2, 529,
	530, 5, 195, 98, 2, 530, 531, 5, 165, 83, 2, 531, 532, 5, 175, 88, 2, 532,
	130, 3, 2, 2, 2, 533, 534, 5, 199, 100, 2, 534, 535, 5, 189, 95, 2, 535,
	132, 3, 2, 2, 2, 536, 537, 5, 165, 83, 2, 537, 538, 5, 201, 101, 2, 538,
	539, 5, 195, 98, 2, 539, 540, 5, 195, 98, 2, 540, 541, 5, 169, 85, 2, 541,
	542, 5, 187, 94, 2, 542, 543, 5, 199, 100, 2, 543, 134, 3, 2, 2, 2, 544,
	545, 5, 187, 94, 2, 545, 546, 5, 169, 85, 2, 546, 547, 5, 205, 103, 2,
	547, 136, 3, 2, 2, 2, 548, 549, 5, 189, 95, 2, 549, 550, 5, 183, 92, 2,
	550, 551, 5, 167, 84, 2, 551, 138, 3, 2, 2, 2, 552, 556, 9, 2, 2, 2, 553,
	555, 9, 3, 2, 2, 554, 553, 3, 2, 2, 2, 555, 558, 3, 2, 2, 2, 556, 554,
	3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 557, 140, 3, 2, 2, 2, 558, 556, 3, 2,
	2, 2, 559, 563, 9, 4, 2, 2, 560, 562, 5, 159, 80, 2, 561, 560, 3, 2, 2,
	2, 562, 565, 3, 2, 2, 2, 563, 561, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2, 564,
	584, 3, 2, 2, 2, 565, 563, 3, 2, 2, 2, 566, 584, 7, 50, 2, 2, 567, 568,
	7, 50, 2, 2, 568, 569, 7, 122, 2, 2, 569, 571, 3, 2, 2, 2, 570, 572, 5,
	157, 79, 2, 571, 570, 3, 2, 2, 2, 572, 573, 3, 2, 2, 2, 573, 571, 3, 2,
	2, 2, 573, 574, 3, 2, 2, 2, 574, 584, 3, 2, 2, 2, 575, 576, 7, 50, 2, 2,
	576, 577, 7, 100, 2, 2, 577, 579, 3, 2, 2, 2, 578, 580, 9, 5, 2, 2, 579,
	578, 3, 2, 2, 2, 580, 581, 3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 581, 582,
	3, 2, 2, 2, 582, 584, 3, 2, 2, 2, 583, 559, 3, 2, 2, 2, 583, 566, 3, 2,
	2, 2, 583, 567, 3, 2, 2, 2, 583, 575, 3, 2, 2, 2, 584, 142, 3, 2, 2, 2,
	585, 589, 9, 4, 2, 2, 586, 588, 5, 159, 80, 2, 587, 586, 3, 2, 2, 2, 588,
	591, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 594,
	3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 592, 594, 7, 50, 2, 2, 593, 585, 3, 2,
	2, 2, 593, 592, 3, 2, 2, 2, 593, 594, 3, 2, 2, 2, 594, 595, 3, 2, 2, 2,
	595, 597, 7, 48, 2, 2, 596, 598, 5, 159, 80, 2, 597, 596, 3, 2, 2, 2, 598,
	599, 3, 2, 2, 2, 599, 597, 3, 2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 610,
	3, 2, 2, 2, 601, 603, 5, 169, 85, 2, 602, 604, 9, 6, 2, 2, 603, 602, 3,
	2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 606, 3, 2, 2, 2, 605, 607, 5, 159,
	80, 2, 606, 605, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608, 606, 3, 2, 2, 2,
	608, 609, 3, 2, 2, 2, 609, 611, 3, 2, 2, 2, 610, 601, 3, 2, 2, 2, 610,
	611, 3, 2, 2, 2, 611, 144, 3, 2, 2, 2, 612, 613, 7, 66, 2, 2, 613, 614,
	5, 139, 70, 2, 614, 146, 3, 2, 2, 2, 615, 623, 7, 41, 2, 2, 616, 617, 7,
	94, 2, 2, 617, 622, 11, 2, 2, 2, 618, 619, 7, 41, 2, 2, 619, 622, 7, 41,
	2, 2, 620, 622, 10, 7, 2, 2, 621, 616, 3, 2, 2, 2, 621, 618, 3, 2, 2, 2,
	621, 620, 3, 2, 2, 2, 622, 625, 3, 2, 2, 2, 623, 621, 3, 2, 2, 2, 623,
	624, 3, 2, 2, 2, 624, 626, 3, 2, 2, 2, 625, 623, 3, 2, 2, 2, 626, 640,
	7, 41, 2, 2, 627, 635, 7, 36, 2, 2, 628, 629, 7, 94, 2, 2, 629, 634, 11,
	2, 2, 2, 630, 631, 7, 36, 2, 2, 631, 634, 7, 36, 2, 2, 632, 634, 10, 8,
	2, 2, 633, 628, 3, 2, 2, 2, 633, 630, 3, 2, 2, 2, 633, 632, 3, 2, 2, 2,
	634, 637, 3, 2, 2, 2, 635, 633, 3, 2, 2, 2, 635, 636, 3, 2, 2, 2, 636,
	638, 3, 2, 2, 2, 637, 635, 3, 2, 2, 2, 638, 640, 7, 36, 2, 2, 639, 615,
	3, 2, 2, 2, 639, 627, 3, 2, 2, 2, 640, 148, 3, 2, 2, 2, 641, 642, 7, 49,
	2, 2, 642, 643, 7, 49, 2, 2, 643, 647, 3, 2, 2, 2, 644, 646, 10, 9, 2,
	2, 645, 644, 3, 2, 2, 2, 646, 649, 3, 2, 2, 2, 647, 645, 3, 2, 2, 2, 647,
	648, 3, 2, 2, 2, 648, 655, 3, 2, 2, 2, 649, 647, 3, 2, 2, 2, 650, 652,
	7, 15, 2, 2, 651, 650, 3, 2, 2, 2, 651, 652, 3, 2, 2, 2, 652, 653, 3, 2,
	2, 2, 653, 656, 7, 12, 2, 2, 654, 656, 7, 2, 2, 3, 655, 651, 3, 2, 2, 2,
	655, 654, 3, 2, 2, 2, 656, 657, 3, 2, 2, 2, 657, 658, 8, 75, 2, 2, 658,
	150, 3, 2, 2, 2, 659, 660, 7, 49, 2, 2, 660, 661, 7, 44, 2, 2, 661, 665,
	3, 2, 2, 2, 662, 664, 11, 2, 2, 2, 663, 662, 3, 2, 2, 2, 664, 667, 3, 2,
	2, 2, 665, 666, 3, 2, 2, 2, 665, 663, 3, 2, 2, 2, 666, 668, 3, 2, 2, 2,
	667, 665, 3, 2, 2, 2, 668, 669, 7, 44, 2, 2, 669, 670, 7, 49, 2, 2, 670,
	671, 3, 2, 2, 2, 671, 672, 8, 76, 2, 2, 672, 152, 3, 2, 2, 2, 673, 674,
	9, 10, 2, 2, 674, 675, 3, 2, 2, 2, 675, 676, 8, 77, 2, 2, 676, 154, 3,
	2, 2, 2, 677, 678, 11, 2, 2, 2, 678, 156, 3, 2, 2, 2, 679, 680, 9, 11,
	2, 2, 680, 158, 3, 2, 2, 2, 681, 682, 9, 12, 2, 2, 682, 160, 3, 2, 2, 2,
	683, 684, 9, 13, 2, 2, 684, 162, 3, 2, 2, 2, 685, 686, 9, 14, 2, 2, 686,
	164, 3, 2, 2, 2, 687, 688, 9, 15, 2, 2, 688, 166, 3, 2, 2, 2, 689, 690,
	9, 16, 2, 2, 690, 168, 3, 2, 2, 2, 691, 692, 9, 17, 2, 2, 692, 170, 3,
	2, 2, 2, 693, 694, 9, 18, 2, 2, 694, 172, 3, 2, 2, 2, 695, 696, 9, 19,
	2, 2, 696, 174, 3, 2, 2, 2, 697, 698, 9, 20, 2, 2, 698, 176, 3, 2, 2, 2,
	699, 700, 9, 21, 2, 2, 700, 178, 3, 2, 2, 2, 701, 702, 9, 22, 2, 2, 702,
	180, 3, 2, 2, 2, 703, 704, 9, 23, 2, 2, 704, 182, 3, 2, 2, 2, 705, 706,
	9, 24, 2, 2, 706, 184, 3, 2, 2, 2, 707, 708, 9, 25, 2, 2, 708, 186, 3,
	2, 2, 2, 709, 710, 9, 26, 2, 2, 710, 188, 3, 2, 2, 2, 711, 712, 9, 27,
	2, 2, 712, 190, 3, 2, 2, 2, 713, 714, 9, 28, 2, 2, 714, 192, 3, 2, 2, 2,
	715, 716, 9, 29, 2, 2, 716, 194, 3, 2, 2, 2, 717, 718, 9, 30, 2, 2, 718,
	196, 3, 2, 2, 2, 719, 720, 9, 31, 2, 2, 720, 198, 3, 2, 2, 2, 721, 722,
	9, 32, 2, 2, 722, 200, 3, 2, 2, 2, 723, 724, 9, 33, 2, 2, 724, 202, 3,
	2, 2, 2, 725, 726, 9, 34, 2, 2, 726, 204, 3, 2, 2, 2, 727, 728, 9, 35,
	2, 2, 728, 206, 3, 2, 2, 2, 729, 730, 9, 36, 2, 2, 730, 208, 3, 2, 2, 2,
	731, 732, 9, 37, 2, 2, 732, 210, 3, 2, 2, 2, 733, 734, 9, 38, 2, 2, 734,
	212, 3, 2, 2, 2, 735, 736, 11, 2, 2, 2, 736, 737, 3, 2, 2, 2, 737, 738,
	8, 107, 3, 2, 738, 214, 3, 2, 2, 2, 26, 2, 293, 413, 425, 556, 563, 573,
	581, 583, 589, 593, 599, 603, 608, 610, 621, 623, 633, 635, 639, 647, 651,
	655, 665, 4, 2, 3, 2, 2, 4, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "ERRORCHANNEL",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'=~'", "'!~'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'?'", "':'", "'::'", "'..'", "','",
	"'('", "')'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "DOT", "T_REGEX_MATCH", "T_REGEX_NON_MATCH", "T_EQ", "T_NE", "T_LT",
	"T_GT", "T_LE", "T_GE", "T_PLUS", "T_MINUS", "T_TIMES", "T_DIV", "T_MOD",
	"T_QUESTION", "T_COLON", "T_SCOPE", "T_RANGE", "T_COMMA", "T_OPEN", "T_CLOSE",
	"T_OBJECT_OPEN", "T_OBJECT_CLOSE", "T_ARRAY_OPEN", "T_ARRAY_CLOSE", "T_AGGREGATE",
	"T_ALL", "T_AND", "T_ANY", "T_ASC", "T_COLLECT", "T_DESC", "T_DISTINCT",
	"T_FALSE", "T_FILTER", "T_FOR", "T_GRAPH", "T_IN", "T_INBOUND", "T_INSERT",
	"T_INTO", "T_K_SHORTEST_PATHS", "T_LET", "T_LIKE", "T_LIMIT", "T_NONE",
	"T_NOT", "T_NULL", "T_OR", "T_OUTBOUND", "T_REMOVE", "T_REPLACE", "T_RETURN",
	"T_SHORTEST_PATH", "T_SORT", "T_TRUE", "T_UPDATE", "T_UPSERT", "T_WITH",
	"T_KEEP", "T_COUNT", "T_OPTIONS", "T_PRUNE", "T_SEARCH", "T_TO", "T_CURRENT",
	"T_NEW", "T_OLD", "T_STRING", "T_INT", "T_FLOAT", "T_PARAMETER", "T_QUOTED_STRING",
	"SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	"ERROR_RECONGNIGION",
}

var lexerRuleNames = []string{
	"DOT", "T_REGEX_MATCH", "T_REGEX_NON_MATCH", "T_EQ", "T_NE", "T_LT", "T_GT",
	"T_LE", "T_GE", "T_PLUS", "T_MINUS", "T_TIMES", "T_DIV", "T_MOD", "T_QUESTION",
	"T_COLON", "T_SCOPE", "T_RANGE", "T_COMMA", "T_OPEN", "T_CLOSE", "T_OBJECT_OPEN",
	"T_OBJECT_CLOSE", "T_ARRAY_OPEN", "T_ARRAY_CLOSE", "T_AGGREGATE", "T_ALL",
	"T_AND", "T_ANY", "T_ASC", "T_COLLECT", "T_DESC", "T_DISTINCT", "T_FALSE",
	"T_FILTER", "T_FOR", "T_GRAPH", "T_IN", "T_INBOUND", "T_INSERT", "T_INTO",
	"T_K_SHORTEST_PATHS", "T_LET", "T_LIKE", "T_LIMIT", "T_NONE", "T_NOT",
	"T_NULL", "T_OR", "T_OUTBOUND", "T_REMOVE", "T_REPLACE", "T_RETURN", "T_SHORTEST_PATH",
	"T_SORT", "T_TRUE", "T_UPDATE", "T_UPSERT", "T_WITH", "T_KEEP", "T_COUNT",
	"T_OPTIONS", "T_PRUNE", "T_SEARCH", "T_TO", "T_CURRENT", "T_NEW", "T_OLD",
	"T_STRING", "T_INT", "T_FLOAT", "T_PARAMETER", "T_QUOTED_STRING", "SINGLE_LINE_COMMENT",
	"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR", "HEX_DIGIT", "DIGIT",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ERROR_RECONGNIGION",
}

type CAQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCAQLLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CAQLLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCAQLLexer(input antlr.CharStream) *CAQLLexer {
	l := new(CAQLLexer)
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
	l.GrammarFileName = "CAQLLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CAQLLexer tokens.
const (
	CAQLLexerDOT                 = 1
	CAQLLexerT_REGEX_MATCH       = 2
	CAQLLexerT_REGEX_NON_MATCH   = 3
	CAQLLexerT_EQ                = 4
	CAQLLexerT_NE                = 5
	CAQLLexerT_LT                = 6
	CAQLLexerT_GT                = 7
	CAQLLexerT_LE                = 8
	CAQLLexerT_GE                = 9
	CAQLLexerT_PLUS              = 10
	CAQLLexerT_MINUS             = 11
	CAQLLexerT_TIMES             = 12
	CAQLLexerT_DIV               = 13
	CAQLLexerT_MOD               = 14
	CAQLLexerT_QUESTION          = 15
	CAQLLexerT_COLON             = 16
	CAQLLexerT_SCOPE             = 17
	CAQLLexerT_RANGE             = 18
	CAQLLexerT_COMMA             = 19
	CAQLLexerT_OPEN              = 20
	CAQLLexerT_CLOSE             = 21
	CAQLLexerT_OBJECT_OPEN       = 22
	CAQLLexerT_OBJECT_CLOSE      = 23
	CAQLLexerT_ARRAY_OPEN        = 24
	CAQLLexerT_ARRAY_CLOSE       = 25
	CAQLLexerT_AGGREGATE         = 26
	CAQLLexerT_ALL               = 27
	CAQLLexerT_AND               = 28
	CAQLLexerT_ANY               = 29
	CAQLLexerT_ASC               = 30
	CAQLLexerT_COLLECT           = 31
	CAQLLexerT_DESC              = 32
	CAQLLexerT_DISTINCT          = 33
	CAQLLexerT_FALSE             = 34
	CAQLLexerT_FILTER            = 35
	CAQLLexerT_FOR               = 36
	CAQLLexerT_GRAPH             = 37
	CAQLLexerT_IN                = 38
	CAQLLexerT_INBOUND           = 39
	CAQLLexerT_INSERT            = 40
	CAQLLexerT_INTO              = 41
	CAQLLexerT_K_SHORTEST_PATHS  = 42
	CAQLLexerT_LET               = 43
	CAQLLexerT_LIKE              = 44
	CAQLLexerT_LIMIT             = 45
	CAQLLexerT_NONE              = 46
	CAQLLexerT_NOT               = 47
	CAQLLexerT_NULL              = 48
	CAQLLexerT_OR                = 49
	CAQLLexerT_OUTBOUND          = 50
	CAQLLexerT_REMOVE            = 51
	CAQLLexerT_REPLACE           = 52
	CAQLLexerT_RETURN            = 53
	CAQLLexerT_SHORTEST_PATH     = 54
	CAQLLexerT_SORT              = 55
	CAQLLexerT_TRUE              = 56
	CAQLLexerT_UPDATE            = 57
	CAQLLexerT_UPSERT            = 58
	CAQLLexerT_WITH              = 59
	CAQLLexerT_KEEP              = 60
	CAQLLexerT_COUNT             = 61
	CAQLLexerT_OPTIONS           = 62
	CAQLLexerT_PRUNE             = 63
	CAQLLexerT_SEARCH            = 64
	CAQLLexerT_TO                = 65
	CAQLLexerT_CURRENT           = 66
	CAQLLexerT_NEW               = 67
	CAQLLexerT_OLD               = 68
	CAQLLexerT_STRING            = 69
	CAQLLexerT_INT               = 70
	CAQLLexerT_FLOAT             = 71
	CAQLLexerT_PARAMETER         = 72
	CAQLLexerT_QUOTED_STRING     = 73
	CAQLLexerSINGLE_LINE_COMMENT = 74
	CAQLLexerMULTILINE_COMMENT   = 75
	CAQLLexerSPACES              = 76
	CAQLLexerUNEXPECTED_CHAR     = 77
	CAQLLexerERROR_RECONGNIGION  = 78
)

// CAQLLexerERRORCHANNEL is the CAQLLexer channel.
const CAQLLexerERRORCHANNEL = 2
