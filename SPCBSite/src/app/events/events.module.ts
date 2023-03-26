import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EventsComponent } from './events.component';
import { MatCardModule } from '@angular/material/card';

@NgModule({
    declarations: [
        MatCardModule,
        EventsComponent
    ],
    imports: [
        MatCardModule,
        CommonModule
    ], 
    exports: [
        MatCardModule,
        EventsComponent
    ]
  })
  export class CaraouselModule { }