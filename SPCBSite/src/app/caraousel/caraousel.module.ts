import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CaraouselComponent } from './caraousel.component';

@NgModule({
  declarations: [
    CaraouselComponent
  ],
  imports: [
    CommonModule
  ], 
  exports: [
    CaraouselComponent
  ]
})
export class CaraouselModule { }
