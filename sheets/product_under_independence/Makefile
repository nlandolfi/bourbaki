# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_under_independence.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

product_under_independence.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > product_under_independence.tex

product_under_independence.pdf: ../../*.tex ../../trademark.pdf *.tex product_under_independence.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_under_independence.tex
	../../latexrun product_under_independence.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_under_independence.pdf
	open product_under_independence.pdf

o:
	make open

terms: product_under_independence.tex
	bbk sheet -mode ts product_under_independence.tex

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
	rm -f product_under_independence.*
	make reset
	bbk sheet -mode mk > Makefile