import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { OfficersComponent } from './officers/officers.component';
import { RentComponent } from './rent/rent.component';
import { SignupComponent } from './signup/signup.component';
import { CaraouselModule } from './caraousel/caraousel.module';

export const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' }, // set default path to 'home'
  { path: 'home', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'officers', component: OfficersComponent },
  { path: 'rent', component: RentComponent },
  { path: 'signup', component: SignupComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
