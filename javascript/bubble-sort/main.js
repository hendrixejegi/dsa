import { duration } from '../utils/duration.js';

function bubbleSort(arr) {
  let length = arr.length - 1;
  let swapped = false;

  do {
    swapped = false;
    for (let i = 0; i < length; i++) {
      if (arr[i] > arr[i + 1]) {
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
