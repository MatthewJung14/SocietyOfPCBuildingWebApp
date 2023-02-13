import { Component } from '@angular/core';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  images = [
    {
      imageSrc:
      '../../assets/IMG_0016.JPG',
      imageAlt: 'GBMpic#16',
    },
    {
      imageSrc:
      '../../assets/IMG_0015.JPG',
      imageAlt: 'GBMpic#15',
    },
    {
      imageSrc:
      '../../assets/IMG_0012.JPG',
      imageAlt: 'GBMpic#12',
    },
    {
      imageSrc:
      '../../assets/IMG_0017.JPG',
      imageAlt: 'GBMpic#17',
    },
  ]
}
