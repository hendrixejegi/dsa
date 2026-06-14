import { createArr } from '../utils/array.js';
import { duration } from '../utils/duration.js';

function selectionSort(arr) {
  for (let i = 0; i < arr.length - 1; i++) {
    let minValueIdx = i;
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[j] < arr[minValueIdx]) {
        minValueIdx = j;
      }
    }
    if (i != minValueIdx) {
      [arr[i], arr[minValueIdx]] = [arr[minValueIdx], arr[i]];
    }
  }
}

(function main() {
  const arr = createArr(10);
  duration(() => selectionSort(arr));
  console.log(arr);
})();
