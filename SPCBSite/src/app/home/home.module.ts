import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import { MatCardModule } from '@angular/material/card';
import { CaraouselModule } from '../caraousel/caraousel.module';

@NgModule({
  declarations: [
    HomeComponent
  ],
  imports: [
    CommonModule,
    MatCardModule,
    CaraouselModule
  ],
  exports: [
    HomeComponent
  ]
})
export class HomeModule { }