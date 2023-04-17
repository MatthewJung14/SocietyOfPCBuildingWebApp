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

  constructor() {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.timeSlot = "";
  }
}
