import { waitForAsync, TestBed, ComponentFixture } from "@angular/core/testing";
import { RouterTestingModule } from "@angular/router/testing";
import { Router } from "@angular/router";
import { Location } from '@angular/common';
import { routes } from "./app-routing.module";
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { OfficersComponent } from './officers/officers.component';
import { RentComponent } from './rent/rent.component';
import { SignupComponent } from './signup/signup.component';

describe("App Routing", ()=>{
    let router: Router;
    let fixture: ComponentFixture<AppComponent>;
    let location: Location;

    beforeEach(waitForAsync(()=>{
      TestBed.configureTestingModule({
        imports: [RouterTestingModule.withRoutes(routes)],
        declarations: [
            AppComponent, 
            HomeComponent,
            LoginComponent,
            OfficersComponent,
            RentComponent,
            SignupComponent
        ]
      }).compileComponents();
    }));

    beforeEach(()=>{
       router = TestBed.inject(Router);
       location = TestBed.inject(Location);
       router.initialNavigation();
       fixture = TestBed.createComponent(AppComponent);
    })

    it("should navigate to default path = home", waitForAsync(() => {
        fixture.detectChanges();
        fixture.whenStable().then(() => {
          expect(location.path()).toBe('/home');
        })
    }));
});