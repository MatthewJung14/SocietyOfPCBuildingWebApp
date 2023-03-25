import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})

@Injectable({
  providedIn: 'root'
})

export class SettingsComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor(private http: HttpClient) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }

  async updateAccount() {
    const headers = new Headers( {
        'Content-Type': 'application/json'
    });
    const data = {FirstName: this.firstName, LastName: this.lastName, Email: this.email, Password: this.password}
    console.log(data);
    this.http.post('http://localhost:4200/api/update-account', data).toPromise();
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
    return;
  }
}

export interface settingsFields {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}