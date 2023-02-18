# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_complements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

set_complements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_complements.tex

set_complements.pdf: ../../*.tex ../../trademark.pdf *.tex set_complements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_complements.tex
	../../latexrun set_complements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_complements.pdf
	open set_complements.pdf

o:
	make open

terms: set_complements.tex
	bbk_sheet -mode ts set_complements.tex

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
	rm -f set_complements.*
	make reset
	bbk_sheet -mode mk > Makefile