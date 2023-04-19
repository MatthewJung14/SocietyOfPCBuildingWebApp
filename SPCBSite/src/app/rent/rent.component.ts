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
  reservedSlots: string[] = []; // Declare an array to store reserved time slots

  constructor() {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.timeSlot = "";

    // Get the stored strings array from local storage
    const storedStrings = localStorage.getItem('strings');
    if (storedStrings) {
      this.strings = JSON.parse(storedStrings);
    }
  }

  submitTime() {
    // Check if the selected time slot is in the strings array
    const index = this.strings.indexOf(this.timeSlot);
  
    if (index !== -1) {
      // If the selected time slot is in the array, check if the email is already in the reserved slots
      const reservedSlotIndex = this.reservedSlots.indexOf(this.email);
      if (reservedSlotIndex !== -1) {
        console.log(`Email ${this.email} has already reserved a time slot.`);
        return;
      }
  
      // If the email is not in the reserved slots, add the email to the reserved slots array and remove the selected time slot from the strings array
      this.reservedSlots.push(this.email);
      this.strings.splice(index, 1);
  
      // Store the updated arrays in local storage
      localStorage.setItem('strings', JSON.stringify(this.strings));
      localStorage.setItem('reservedSlots', JSON.stringify(this.reservedSlots));
    }
  }  
}
