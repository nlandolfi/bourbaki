# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make variation_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

variation_measure.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > variation_measure.tex

variation_measure.pdf: ../../*.tex ../../trademark.pdf *.tex variation_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode variation_measure.tex
	../../latexrun variation_measure.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: variation_measure.pdf
	open variation_measure.pdf

o:
	make open

terms: variation_measure.tex
	bbk sheet -mode ts variation_measure.tex

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
	rm -f variation_measure.*
	make reset
	bbk sheet -mode mk > Makefile