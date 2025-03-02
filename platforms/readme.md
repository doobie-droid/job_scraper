### COMPREHENSIVE LIST OF ALL PLATFORMS USED IN BUILDING THIS TOOL

1. [Linkedin](https://linkedin.com) -> data from linkedin is pooled using rapid api's api and also using a dedicated web scraper to search using the keywords listed in the .env
2. [TestGorilla](https://www.testgorilla.com) -> data from testGorilla is pooled using a web scraper checking out their career page
3. [Workable](https://www.workable.com/) -> data from workable is pooled using a web scraper checking out their jobs page
4. [RemoteAfrica](https://remoteafrica.io/) -> data from remoteAfrica is pooled using a web scraper checking out their jobs page
5. [GolangProjects](https://www.golangprojects.com/) -> data from Golang Projects is pooled using a scraper. It only contains golang jobs and it only contains remote or non-remote type jobs
6. [BreezyHr](https://breezy.hr/) -> This uses googles api to cumulate a list of of all the companies that use breezy.hr, then uses a web crawler to visit all those company pages

### TODO I would love to integrate the following Platforms

- [Lever](https://www.lever.co/) -> You could go the route of using google apis to cumulate a list of all the companies that use lever.co and crawl-visit each of those company pages
- [WellFound](https://wellfound.com) -> This already has a way of filtering available jobs on their website, so you could just visit their website (Not a priority though as not a lot of worldwide jobs there)
- [AshbyHq](https://ashbyhq.com) -> You could try the google api route then crawl cumulated list
  site:bamboo.hr
