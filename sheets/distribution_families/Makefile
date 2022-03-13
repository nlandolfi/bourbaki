# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distribution_families.pdf

distribution_families.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_families.tex

distribution_families.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_families.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distribution_families.tex
	../../latexrun distribution_families.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_families.pdf
	open distribution_families.pdf

o:
	make open

terms: distribution_families.tex
	bbk_sheet -mode ts distribution_families.tex

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
	rm -f distribution_families.*
	make reset
	bbk_sheet -mode mk > Makefile