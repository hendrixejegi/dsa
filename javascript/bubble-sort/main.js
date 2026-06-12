import { duration } from '../utils/duration.js';

function bubbleSort(arr) {
  let length = arr.length;
  let swapped;

  do {
    swapped = false;
    for (let i = 0; i < length - 1; i++) {
      let currentElem = arr[i];
      let nextElem = arr[i + 1];

      if (currentElem > nextElem) {
        // arr[i] = nextElem;
        // arr[i + 1] = currentElem;

        // In-place swap
        [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];

        swapped = true;
      }
    }
    length--;
  } while (swapped);
}

(function main() {
  const arr = [];

  // 300 is the length of the array
  for (let i = 0; i < 300; i++) {
    arr[i] = Math.round(Math.random() * 1000);
  }

  duration(() => {
    bubbleSort(arr);
    console.log(arr);
  });
})();
