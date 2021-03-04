\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{bsheet}[2021/03/03 A simple class file for sheets]
% choose here
\def\FONT{6}


\if\FONT1
% option 1, cmss tex and cmmi math
\renewcommand\itdefault{sl}
\renewcommand{\rmdefault}{cmss}
\fi
\if\FONT2
% option 2, concrete roman and concrete math
\usepackage{concmath}
\fi
\if\FONT3
% option 3, cmbright
\usepackage{cmbright}
\providecommand{\iestyle}[1]{{\fontshape{it}\selectfont #1}}
\fi
\if\FONT4
% option 4, concrete roman and euler math
\usepackage{beton}
\usepackage[euler-digits,euler-hat-accent]{eulervm}
\fi
\if\FONT5
% option 5, cmss tex and euler math
\renewcommand\itdefault{sl}
\renewcommand{\rmdefault}{cmss}
\usepackage[euler-digits,euler-hat-accent]{eulervm}
\fi
\if\FONT6
% option 6, cmss tex and concrete math
\fi
