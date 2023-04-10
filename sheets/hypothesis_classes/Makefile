# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make hypothesis_classes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

hypothesis_classes.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > hypothesis_classes.tex

hypothesis_classes.pdf: ../../*.tex ../../trademark.pdf *.tex hypothesis_classes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode hypothesis_classes.tex
	../../latexrun hypothesis_classes.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: hypothesis_classes.pdf
	open hypothesis_classes.pdf

o:
	make open

terms: hypothesis_classes.tex
	bbk sheet -mode ts hypothesis_classes.tex

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
	rm -f hypothesis_classes.*
	make reset
	bbk sheet -mode mk > Makefile