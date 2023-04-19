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

  strings: string[] = ['TIMES AVAILABLE', '8AM', '9AM', '10AM', '11AM', '12PM', '1PM', '2PM', '3PM', '4PM', '5PM', '6PM', '7PM', '8PM', '9PM', '10PM']; // Define the array of strings

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
