# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make constrained_quadratic_form_extrema.pdf

constrained_quadratic_form_extrema.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > constrained_quadratic_form_extrema.tex

constrained_quadratic_form_extrema.pdf: ../../*.tex ../../trademark.pdf *.tex constrained_quadratic_form_extrema.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode constrained_quadratic_form_extrema.tex
	../../latexrun constrained_quadratic_form_extrema.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: constrained_quadratic_form_extrema.pdf
	open constrained_quadratic_form_extrema.pdf

o:
	make open

terms: constrained_quadratic_form_extrema.tex
	bbk_sheet -mode ts constrained_quadratic_form_extrema.tex

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
	rm -f constrained_quadratic_form_extrema.*
	make reset
	bbk_sheet -mode mk > Makefile