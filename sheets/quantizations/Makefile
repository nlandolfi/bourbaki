# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quantizations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

quantizations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > quantizations.tex

quantizations.pdf: ../../*.tex ../../trademark.pdf *.tex quantizations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quantizations.tex
	../../latexrun quantizations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quantizations.pdf
	open quantizations.pdf

o:
	make open

terms: quantizations.tex
	bbk_sheet -mode ts quantizations.tex

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
	rm -f quantizations.*
	make reset
	bbk_sheet -mode mk > Makefile