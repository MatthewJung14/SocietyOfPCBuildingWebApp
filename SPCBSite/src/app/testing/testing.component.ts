import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-testing',
  templateUrl: './testing.component.html',
  styleUrls: ['./testing.component.css']
})
export class TestingComponent {
  firstName: string;
  lastName: string;
  email: string;
  password: string;

  constructor(private http: HttpClient, public authService: AuthService) {
    this.firstName = "";
    this.lastName = "";
    this.email = "";
    this.password = "";
  }

  async login() {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    });
    const data = { Email: this.email, Password: this.password };
    console.log(data);
    await this.http.post('http://localhost:4200/api/login', data, { headers }).toPromise();
    this.email = "";
    this.password = "";
    return;
  }
}
