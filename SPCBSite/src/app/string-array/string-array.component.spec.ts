import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StringArrayComponent } from './string-array.component';

describe('StringArrayComponent', () => {
  let component: StringArrayComponent;
  let fixture: ComponentFixture<StringArrayComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StringArrayComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StringArrayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
