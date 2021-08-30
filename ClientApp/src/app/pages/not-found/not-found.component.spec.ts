import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatButtonModule } from '@angular/material/button';
import { RouterTestingModule } from '@angular/router/testing';
import { RouterLinkWithHref } from '@angular/router';

import { NotFoundComponent } from './not-found.component';
import { TestingPage } from 'src/app/testing';

class NotFoundComponentTestingPage extends TestingPage {
  get h1() { return this.query<HTMLHeadingElement>('h1'); }
  get h2() { return this.query<HTMLHeadingElement>('h2'); }
  get link() { return this.query<HTMLAnchorElement>('a'); }
  get routerLink() { 
    return this.queryByCss('a').injector.get(RouterLinkWithHref); 
  }

  constructor(protected fixture: ComponentFixture<NotFoundComponent>) {
    super(fixture);
  }
}

describe('NotFoundComponent', () => {
  let component: NotFoundComponent;
  let fixture: ComponentFixture<NotFoundComponent>;
  let page: NotFoundComponentTestingPage;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NotFoundComponent ],
      imports: [
        RouterTestingModule.withRoutes([]),
        MatButtonModule
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NotFoundComponent);
    component = fixture.componentInstance;
    page = new NotFoundComponentTestingPage(fixture);
    fixture.detectChanges();
  });

  describe('#template', () => {
    it('should contain a h1 and h2 element with error information, and an anchor element', () => {
      expect(page.h1).toBeTruthy();
      expect(page.h1.textContent).toContain('404');
      expect(page.h2).toBeTruthy();
      expect(page.h2.textContent).toContain('Oops! The page you were looking for was moved or doesn\'t exist');
      expect(page.link).toBeTruthy();
      expect(page.link.textContent).toContain('Back Home');
    });

    it('routerLink should be setup to redirect users to `/`', () => {
      expect(page.routerLink['commands']).toEqual(['/']);
    });
  });
});
