import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
 public name: string = '';
 constructor(private route: ActivatedRoute, private router: Router) {}

 ngOnInit(): void {
  this.route.queryParams.subscribe((queryParam) => {
    this.name = queryParam['name'];
    console.log(queryParam);
  })
 }

 goToSignup(): void {
   this.router.navigate(['/app/login']);
 }
}
