import { createArr } from '../../utils/array.js';
import { duration } from '../../utils/duration.js';

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
  const arr = createArr(300);
  duration(() => bubbleSort(arr));
  console.log(arr);
})();
