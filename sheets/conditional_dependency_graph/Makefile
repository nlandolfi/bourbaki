conditional_dependency_graph.tex: sheet.tex
	sheet -mode c -in sheet.tex > conditional_dependency_graph.tex

conditional_dependency_graph.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_dependency_graph.tex
	pdflatex conditional_dependency_graph.tex
	make defs

make:
	make conditional_dependency_graph.pdf

open: conditional_dependency_graph.pdf
	open conditional_dependency_graph.pdf

o:
	make open

defs: conditional_dependency_graph.tex
	defs conditional_dependency_graph.tex

spell:
	aspell -c conditional_dependency_graph.tex

s:
	make spell