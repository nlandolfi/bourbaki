# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make distribution_approximation.pdf
	make graph.pdf

distribution_approximation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_approximation.tex

distribution_approximation.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_approximation.tex
	# pdflatex --file-line-error -interaction=nonstopmode distribution_approximation.tex
	../../latexrun distribution_approximation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_approximation.pdf
	open distribution_approximation.pdf

o:
	make open

terms: distribution_approximation.tex
	bbk_sheet -mode ts distribution_approximation.tex

spell:
	aspell -c distribution_approximation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f distribution_approximation.*
	make reset
	bbk_sheet -mode mk > Makefile