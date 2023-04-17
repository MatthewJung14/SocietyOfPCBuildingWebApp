import { Component } from '@angular/core';
import { StringArrayComponent } from '../string-array/string-array.component';

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
}
