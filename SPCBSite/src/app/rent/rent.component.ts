import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-rent',
  templateUrl: './rent.component.html',
  styleUrls: ['./rent.component.css']
})
export class RentComponent implements OnInit {
  firstName: string;
  lastName: string;
  email: string;
  timeSlot: string;
  strings: string[];
  emails: string[] = [];

  constructor(private http: HttpClient) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.timeSlot = "";
    this.strings = [];
  }

  ngOnInit(): void {
    this.http.get<string[]>('http://localhost:4200/api/get-event').subscribe((data: string[]) => {
      this.strings = data;
    });
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

      this.firstName = "";
      this.lastName = "";
      this.email = "";
      this.timeSlot = "";
    }
  }
}
