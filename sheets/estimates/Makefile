# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make estimates.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

estimates.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > estimates.tex

estimates.pdf: ../../*.tex ../../trademark.pdf *.tex estimates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode estimates.tex
	../../latexrun estimates.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: estimates.pdf
	open estimates.pdf

o:
	make open

terms: estimates.tex
	bbk_sheet -mode ts estimates.tex

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
	rm -f estimates.*
	make reset
	bbk_sheet -mode mk > Makefile