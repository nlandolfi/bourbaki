# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make polynomial_fit_models.pdf

polynomial_fit_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > polynomial_fit_models.tex

polynomial_fit_models.pdf: ../../*.tex ../../trademark.pdf *.tex polynomial_fit_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode polynomial_fit_models.tex
	../../latexrun polynomial_fit_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: polynomial_fit_models.pdf
	open polynomial_fit_models.pdf

o:
	make open

terms: polynomial_fit_models.tex
	bbk_sheet -mode ts polynomial_fit_models.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f polynomial_fit_models.*
	make reset
	bbk_sheet -mode mk > Makefile