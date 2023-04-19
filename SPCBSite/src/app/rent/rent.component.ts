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
  emails: string[] = [];

  constructor() {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.timeSlot = "";
  }

  submitTime() {
    // Check if the selected time slot is in the strings array
    const index = this.strings.indexOf(this.timeSlot);
  
    if (index !== -1) {
      // If the selected time slot is in the array, check if the email is already in the reserved slots
      const reservedSlots = this.emails.filter(str => str.includes(this.email));
      if (reservedSlots.length > 0) {
        console.log(`Email ${this.email} has already reserved a time slot.`);
        return;
      }
  
      // If the email is not in the reserved slots, add the email to the emails array and remove the time slot from the strings array
      this.emails.push(this.email);
      this.strings.splice(index, 1);
  
      // Store the updated strings array in local storage
      localStorage.setItem('strings', JSON.stringify(this.strings));
    }
  }
}
