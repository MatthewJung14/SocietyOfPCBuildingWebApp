import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core'

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
    console.log("Created this function for cypress.")
    this.change.emit("Home was Clicked")
  }
}
