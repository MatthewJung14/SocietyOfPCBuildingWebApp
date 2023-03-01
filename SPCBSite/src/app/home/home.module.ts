import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import { MatCardModule } from '@angular/material/card';

@NgModule({
    declarations: [
        MatCardModule,
        HomeComponent
    ],
    imports: [
        MatCardModule,
        CommonModule
    ], 
    exports: [
        MatCardModule,
        HomeComponent
    ]
  })
  export class CaraouselModule { }