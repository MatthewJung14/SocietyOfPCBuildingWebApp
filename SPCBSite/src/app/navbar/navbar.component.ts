import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { AuthService } from '../auth.service';
import {LoginFields} from '../login/login.component';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
  providers: [AuthService],
})
export class NavbarComponent {
  loggedIn = false;
  constructor(public authService: AuthService) {}

  ngOnInit(): void {

  }

  @Output() change = new EventEmitter()

  homeClick(): void {
    this.change.emit("Home was Clicked")
  }
}
