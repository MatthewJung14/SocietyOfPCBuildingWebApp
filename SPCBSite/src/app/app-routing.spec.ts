import { waitForAsync, TestBed, ComponentFixture } from "@angular/core/testing";
import { RouterTestingModule } from "@angular/router/testing";
import { Router } from "@angular/router";
import { Location } from '@angular/common';
import { DebugElement } from "@angular/core";
import { routes } from "./app-routing.module";
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { OfficersComponent } from './officers/officers.component';
import { RentComponent } from './rent/rent.component';
import { SignupComponent } from './signup/signup.component';
import { NavbarComponent } from './navbar/navbar.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatCardModule } from '@angular/material/card';
import { CaraouselModule } from './caraousel/caraousel.module';
import { By } from '@angular/platform-browser';

describe("App Routing", ()=>{
    let router: Router;
    let fixture: ComponentFixture<AppComponent>;
    let navFixture: ComponentFixture<NavbarComponent>;
    let location: Location;
    let el: DebugElement;

    beforeEach(waitForAsync(()=>{
      TestBed.configureTestingModule({
        imports: [RouterTestingModule.withRoutes(routes), MatToolbarModule, MatCardModule, CaraouselModule],
        declarations: [
            AppComponent, 
            HomeComponent,
            LoginComponent,
            OfficersComponent,
            RentComponent,
            SignupComponent,
            NavbarComponent
        ],
        providers: [{provide: 'defaultUrl', useValue: '/home'}]
      }).compileComponents();
    }));

    beforeEach(()=>{
       router = TestBed.inject(Router);
       location = TestBed.inject(Location);
       router.initialNavigation();
       fixture = TestBed.createComponent(AppComponent);
       navFixture = TestBed.createComponent(NavbarComponent);
       el = navFixture.debugElement;
    })

    //Defaut path set to home test case
    it("should navigate to default path = home", waitForAsync(() => {
        fixture.detectChanges();
        fixture.whenStable().then(() => {
          expect(location.path()).toBe('/home');
        })
    }));

    //Home Button test case
    it("should navigate to home when home is clicked", waitForAsync(() => {
        navFixture.detectChanges();
        let links = el.queryAll(By.css('button'));
        links[0].nativeElement.click();
        navFixture.whenStable().then(()=>{
            expect(location.path()).toBe('/home');
        })
    }));

    //Officers button test case
    it("should navigate to officers when officers is clicked", waitForAsync(() => {
        navFixture.detectChanges();
        let links = el.queryAll(By.css('button'));
        links[1].nativeElement.click();
        navFixture.whenStable().then(()=>{
            expect(location.path()).toBe('/officers');
        })
    }));

    //Rent a pc button test case
    it("should navigate to rent when rent is clicked", waitForAsync(() => {
        navFixture.detectChanges();
        let links = el.queryAll(By.css('button'));
        links[2].nativeElement.click();
        navFixture.whenStable().then(()=>{
            expect(location.path()).toBe('/rent');
        })
    }));

    //Login button test case
    it("should navigate to login when login is clicked", waitForAsync(() => {
        navFixture.detectChanges();
        let links = el.queryAll(By.css('button'));
        links[3].nativeElement.click();
        navFixture.whenStable().then(()=>{
            expect(location.path()).toBe('/login');
        })
    }));

    //Signup button test case
    it("should navigate to signup when signup is clicked", waitForAsync(() => {
        navFixture.detectChanges();
        let links = el.queryAll(By.css('button'));
        links[4].nativeElement.click();
        navFixture.whenStable().then(()=>{
            expect(location.path()).toBe('/signup');
        })
    }));
});