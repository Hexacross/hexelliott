# http://www.pyimagesearch.com/2016/10/24/ubuntu-16-04-how-to-install-opencv/

# $ mkvirtualenv cv -p python2
# Creates a virtual environment named cv
# Change to python3 if needed.

# $ workon cv
# Verifies you're working in that environment

import cv2
import numpy as np

img = cv2.imread('/home/philosopher/Downloads/armin.png')
cv2.imshow('image', img)
cv2.waitKey(0)
cv2.destroyAllWindows()
