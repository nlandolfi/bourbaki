# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distribution_selectors.pdf

distribution_selectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_selectors.tex

distribution_selectors.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_selectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distribution_selectors.tex
	../../latexrun distribution_selectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_selectors.pdf
	open distribution_selectors.pdf

o:
	make open

terms: distribution_selectors.tex
	bbk_sheet -mode ts distribution_selectors.tex

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
	rm -f distribution_selectors.*
	make reset
	bbk_sheet -mode mk > Makefile