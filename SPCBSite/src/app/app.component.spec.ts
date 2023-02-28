import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButton } from '@angular/material/button';
import { MatCard } from '@angular/material/card';
import { MatFormField } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AppComponent, MatButton, MatCard, MatFormField],
      imports: [MatInputModule, BrowserAnimationsModule],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the app', () => {
    expect(component).toBeTruthy();
  });

  it('should display a login button', () => {
    const button = fixture.nativeElement.querySelector('button');
    expect(button.textContent).toContain('Login');
  });

  it('should display a card with a title', () => {
    const card = fixture.nativeElement.querySelector('mat-card');
    expect(card).toBeTruthy();

    const title = card.querySelector('mat-card-title');
    expect(title).toBeTruthy();
    expect(title.textContent).toContain('Welcome');
  });

  it('should display a form with a username and password field', () => {
    const form = fixture.nativeElement.querySelector('form');
    expect(form).toBeTruthy();

    const usernameInput = form.querySelector('input[formcontrolname="username"]');
    expect(usernameInput).toBeTruthy();

    const passwordInput = form.querySelector('input[formcontrolname="password"]');
    expect(passwordInput).toBeTruthy();
  });
});
