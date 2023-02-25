# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make area_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

area_measure.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > area_measure.tex

area_measure.pdf: ../../*.tex ../../trademark.pdf *.tex area_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode area_measure.tex
	../../latexrun area_measure.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: area_measure.pdf
	open area_measure.pdf

o:
	make open

terms: area_measure.tex
	bbk_sheet -mode ts area_measure.tex

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
	rm -f area_measure.*
	make reset
	bbk_sheet -mode mk > Makefile