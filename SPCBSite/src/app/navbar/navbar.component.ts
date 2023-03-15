import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import {LoginFields} from '../login/login.component';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  constructor() {}

  ngOnInit(): void {}

  @Output() change = new EventEmitter()

  homeClick(): void {
    this.change.emit("Home was Clicked")
  }
}
