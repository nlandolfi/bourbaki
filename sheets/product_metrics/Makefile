# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_metrics.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

product_metrics.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > product_metrics.tex

product_metrics.pdf: ../../*.tex ../../trademark.pdf *.tex product_metrics.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_metrics.tex
	../../latexrun product_metrics.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_metrics.pdf
	open product_metrics.pdf

o:
	make open

terms: product_metrics.tex
	bbk sheet -mode ts product_metrics.tex

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
	rm -f product_metrics.*
	make reset
	bbk sheet -mode mk > Makefile