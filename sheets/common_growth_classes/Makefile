# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make common_growth_classes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

common_growth_classes.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > common_growth_classes.tex

common_growth_classes.pdf: ../../*.tex ../../trademark.pdf *.tex common_growth_classes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode common_growth_classes.tex
	../../latexrun common_growth_classes.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: common_growth_classes.pdf
	open common_growth_classes.pdf

o:
	make open

terms: common_growth_classes.tex
	bbk sheet -mode ts common_growth_classes.tex

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
	rm -f common_growth_classes.*
	make reset
	bbk sheet -mode mk > Makefile