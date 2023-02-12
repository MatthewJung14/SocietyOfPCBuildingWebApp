import { Component, Input, OnInit} from '@angular/core';

interface SPCBImages {
  imageSrc: string;
  imageAlt: string;
}

@Component({
  selector: 'app-caraousel',
  templateUrl: './caraousel.component.html',
  styleUrls: ['./caraousel.component.scss']
})
export class CaraouselComponent implements OnInit{

  @Input() images: SPCBImages[] = []
  @Input() indicators = true;
  @Input() controls = true;
  @Input() autoSlide = false;
  @Input() slideInterval = 8000; //8 seconds

  selectedIndex = 0;

  constructor() { }

  ngOnInit(): void { 
    if(this.autoSlide){
      this.autoSlideImages();
    }
  }

  selectImage(index: number): void {
    this.selectedIndex=index;
  }

  prevClick(): void{
    if (this.selectedIndex === 0){
      this.selectedIndex = this.images.length-1;
    } else {
      this.selectedIndex--;
    }
  }

  nextClick(): void{
    if (this.selectedIndex === this.images.length-1){
      this.selectedIndex = 0;
    } else {
      this.selectedIndex++;
    }
  }

  autoSlideImages(): void {
    setInterval(() =>{
      this.nextClick();
    }, this.slideInterval);
  }
}
