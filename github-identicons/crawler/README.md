# GitHub Identicons

## Observations

- Background: rgb(240, 240, 240)
- Symmetric on the y-axis, so only need to check 3 out of 5 columns
- Dimensions: 420 x 420
    - 35 x 35 white square on top left
    - 34 x 34 white square on bottom right
- To add color to row `i` + column `j`, create a 71 x 71 square centered at image row `70 * i` + column `70 * j`

## Findings

- Username '22353' has an empty iconogram

```
  1. "username"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         27558
	Smallest value:        1
	Largest value:         33,000
	Sum:                   409,265,282
	Mean:                  14,851.052
	Median:                14,069.5
	StDev:                 9,182.179
	Most decimal places:   0
	Most common values:    1 (1x)
	                       2 (1x)
	                       3 (1x)
	                       4 (1x)
	                       5 (1x)

  2. "r1c1"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13785x)
	                       False (13773x)

  3. "r2c1"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13804x)
	                       False (13754x)

  4. "r3c1"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13832x)
	                       False (13726x)

  5. "r4c1"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13846x)
	                       True (13712x)

  6. "r5c1"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13875x)
	                       False (13683x)

  7. "r1c2"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13934x)
	                       False (13624x)

  8. "r2c2"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13871x)
	                       True (13687x)

  9. "r3c2"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13818x)
	                       False (13740x)

 10. "r4c2"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13837x)
	                       False (13721x)

 11. "r5c2"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13789x)
	                       True (13769x)

 12. "r1c3"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13800x)
	                       False (13758x)

 13. "r2c3"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13865x)
	                       True (13693x)

 14. "r3c3"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13815x)
	                       True (13743x)

 15. "r4c3"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    False (13851x)
	                       True (13707x)

 16. "r5c3"

	Type of data:          Boolean
	Contains null values:  False
	Non-null values:       27558
	Unique values:         2
	Most common values:    True (13857x)
	                       False (13701x)

 17. "red"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         168
	Smallest value:        66
	Largest value:         233
	Sum:                   4,562,507
	Mean:                  165.56
	Median:                167
	StDev:                 46.657
	Most decimal places:   0
	Most common values:    217 (482x)
	                       214 (463x)
	                       212 (436x)
	                       213 (436x)
	                       216 (436x)

 18. "green"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         168
	Smallest value:        66
	Largest value:         233
	Sum:                   4,581,905
	Mean:                  166.264
	Median:                169
	StDev:                 46.444
	Most decimal places:   0
	Most common values:    217 (485x)
	                       213 (459x)
	                       219 (455x)
	                       218 (450x)
	                       216 (450x)

 19. "blue"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         167
	Smallest value:        66
	Largest value:         232
	Sum:                   4,566,434
	Mean:                  165.703
	Median:                169
	StDev:                 46.647
	Most decimal places:   0
	Most common values:    215 (496x)
	                       214 (479x)
	                       217 (463x)
	                       213 (458x)
	                       220 (451x)

 20. "hue"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         17150
	Smallest value:        0
	Largest value:         359.545
	Sum:                   4,935,074.627
	Mean:                  179.08
	Median:                178.065
	StDev:                 103.677
	Most decimal places:   6
	Most common values:    120 (60x)
	                       300 (54x)
	                       0 (51x)
	                       180 (50x)
	                       60 (49x)

 21. "saturation"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         1608
	Smallest value:        0.444
	Largest value:         0.656
	Sum:                   15,167.15
	Mean:                  0.55
	Median:                0.551
	StDev:                 0.058
	Most decimal places:   6
	Most common values:    0.5 (394x)
	                       0.6 (320x)
	                       0.556 (175x)
	                       0.636 (144x)
	                       0.455 (140x)

 22. "lightness"

	Type of data:          Number
	Contains null values:  False
	Non-null values:       27558
	Unique values:         104
	Smallest value:        0.549
	Largest value:         0.751
	Sum:                   17,919.033
	Mean:                  0.65
	Median:                0.651
	StDev:                 0.058
	Most decimal places:   6
	Most common values:    0.714 (315x)
	                       0.735 (300x)
	                       0.725 (297x)
	                       0.659 (296x)
	                       0.692 (295x)

Row count: 27558
```