import { Component } from '@angular/core';

@Component({
  selector: 'app-rent',
  templateUrl: './rent.component.html',
  styleUrls: ['./rent.component.css']
})
export class RentComponent {
  firstName: string;
  lastName: string;
  email: string;
  timeSlot: string;

  strings: string[] = ['TIMES AVAILABLE', '10PM', '11PM']; // Define the array of strings

  constructor() {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.timeSlot = "";
  }

  submitTime() {
    // Check if the entered value is present in the array
    const index = this.strings.indexOf(this.timeSlot);
    if (index !== -1) {
      // Remove the element from the array
      this.strings.splice(index, 1);
      console.log(`Removed ${this.timeSlot} from the array.`);
    } else {
      console.log(`${this.timeSlot} not found in the array.`);
    }
  }
  
}
