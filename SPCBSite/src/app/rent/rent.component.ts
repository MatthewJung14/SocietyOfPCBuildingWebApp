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
    // Your code to handle the submission of the form
    console.log('Form submitted!');
  }
}
