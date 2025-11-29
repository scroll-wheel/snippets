#!/usr/bin/env python3

import argparse
import math
import random

BASE = 95
FILE_LENGTH = 80 * 24
N = BASE ** FILE_LENGTH

def to_base95(decimal):
	if decimal == 0:
		return ' '
	reversed = ''
	while decimal != 0:
		remainder = decimal % BASE
		reversed += chr(remainder + 32)
		decimal //= BASE
	return reversed[::-1]

def to_decimal(base95):
	reversed = base95[::-1]
	result = 0
	for i in range(len(reversed)):
		val = ord(reversed[i]) - 32
		result += (val * BASE ** i)
	return result

def generate_keys():
	c = random.randrange(N)

	# Multiplicative inverse exists iff C and N are coprime
	while math.gcd(c, N) != 1:
		c -= 1
	
	i = pow(c, -1, N)
	print(to_base95(c))
	print(to_base95(i))

def main():
	parser = argparse.ArgumentParser(description= \
		'Use mod 95^(80*24) arithmetic to \
		(weakly) encrypt or decrypt standard input')
	mutex_group = parser.add_mutually_exclusive_group()
	mutex_group.add_argument('-g', '--generate', action='store_true', \
		help='generate random multiplicative inverse pair')
	mutex_group.add_argument('-d', '--decrypt', action='store_true', \
		help='decrypt data')
	args = parser.parse_args()

	if args.generate:
		generate_keys()
		return
	
	c, i = None, None
	with open('keys.txt', encoding="utf-8") as f:
		c = to_decimal(f.readline()[:-1])
		i = to_decimal(f.readline()[:-1])
	
	if args.decrypt:
		plaintext = (to_decimal(input()) * i) % N
		print(to_base95(plaintext))
	else:
		ciphertext = (to_decimal(input()) * c) % N
		print(to_base95(ciphertext))

if __name__ == '__main__':
	main()

