import { SimpleChange, SimpleChanges } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestingPage } from 'src/app/testing/testing-page';
import { AvatarComponent } from './avatar.component';

class AvatarComponentTestingPage extends TestingPage {
  get content() { return this.query<HTMLDivElement>('.avatar-content'); }
  get img() { return this.query<HTMLImageElement>('.avatar-image'); }
  get initials() { return this.query<HTMLSpanElement>('.avatar-initials'); }

  constructor(fixture: ComponentFixture<AvatarComponent>) {
    super(fixture);
  }
}

describe('AvatarComponent', () => {
  let component: AvatarComponent;
  let fixture: ComponentFixture<AvatarComponent>;
  let page: AvatarComponentTestingPage;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AvatarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AvatarComponent);
    component = fixture.componentInstance;
    page = new AvatarComponentTestingPage(fixture);
    fixture.detectChanges();
  });

  it('#setInitials() should set #initials to first 1-2 of name param', () => {
    component.setInitials('Alex');
    expect(component.initials).toEqual('A');
    
    component.setInitials('alex deakins');
    expect(component.initials).toEqual('AD');

    component.setInitials('Alex Sean Deakins');
    expect(component.initials).toEqual('AS');
  });

  it('#getFontSize() should return value of 55% size of #diameter with correct units', () => {
    component.diameter = '100';
    expect(component.getFontSize()).toEqual('55px');
    
    component.diameter = '100px';
    expect(component.getFontSize()).toEqual('55px');

    component.diameter = '100%';
    expect(component.getFontSize()).toEqual('55%');

    component.diameter = '10rem';
    expect(component.getFontSize()).toEqual('5rem');

    component.diameter = '10em';
    expect(component.getFontSize()).toEqual('5em');

    component.diameter = '10vh';
    expect(component.getFontSize()).toEqual('5vh');

    component.diameter = '10vw';
    expect(component.getFontSize()).toEqual('5vw');

    component.diameter = '10vmin';
    expect(component.getFontSize()).toEqual('5vmin');

    component.diameter = '10vmax';
    expect(component.getFontSize()).toEqual('5vmax');

    component.diameter = '10cm';
    expect(component.getFontSize()).toEqual('5cm');

    component.diameter = '10mm';
    expect(component.getFontSize()).toEqual('5mm');

    component.diameter = '10in';
    expect(component.getFontSize()).toEqual('5in');

    component.diameter = '10ex';
    expect(component.getFontSize()).toEqual('5ex');

    component.diameter = '10ch';
    expect(component.getFontSize()).toEqual('5ch');
  });

  describe('#ngOnChanges()', () => {
    it('should call #setInitials() when changes contains a falsy avatar and name is truthy', () => {
      const setInitialsSpy = spyOn(component, 'setInitials').and.callFake((name: string) => {});
      const changes: SimpleChanges = {
        avatar: new SimpleChange(null, null, false),
        name: new SimpleChange(null, 'Alex Deakins', false)
      };
      component.ngOnChanges(changes);
      expect(setInitialsSpy).toHaveBeenCalledWith('Alex Deakins');
    });

    it('when diameter changes should set #wrapperStyles height and width to #diameter value and set #fontSize', () => {
      const getFontSizeSpy = spyOn(component, 'getFontSize').and.returnValue('55px');
      const changes: SimpleChanges = {
        diameter: new SimpleChange(null, '100px', false),
      };
      component.diameter = '100px';
      component.ngOnChanges(changes);

      expect(getFontSizeSpy).toHaveBeenCalled();
      expect(component.fontSize).toEqual('55px');
      expect(component.wrapperStyles).toEqual({ height: '100px', width: '100px' });
    });
  });

  describe('#template', () => {
    beforeEach(() => {
      component.wrapperStyles = {
        height: '50px',
        width: '50px',
      };
      component.fontSize = '22.5px';
      component.name = 'Alex Deakins';
      component.initials = 'AD';
      fixture.detectChanges();
    });

    it('should have initials element when #avatar is null', () => {
      component.avatar = null;
      fixture.detectChanges();
      expect(page.img).toBeFalsy();
      expect(page.initials).toBeTruthy();
      expect(page.initials.textContent).toContain('AD');
      expect(page.initials.style.fontSize).toEqual('22.5px');
    });

    it('should have img when #avatar is truthy', () => {
      component.avatar = 'https://example.com/some-user-avatar.png';
      fixture.detectChanges();
      expect(page.initials).toBeFalsy();
      expect(page.img).toBeTruthy();
      expect(page.img.alt).toContain('Alex Deakins avatar');
      expect(page.img.src).toEqual('https://example.com/some-user-avatar.png');
    });
    
  });

});
