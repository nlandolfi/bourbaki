# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_powers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

set_powers.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > set_powers.tex

set_powers.pdf: ../../*.tex ../../trademark.pdf *.tex set_powers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_powers.tex
	../../latexrun set_powers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_powers.pdf
	open set_powers.pdf

o:
	make open

terms: set_powers.tex
	bbk_sheet -mode ts set_powers.tex

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
	rm -f set_powers.*
	make reset
	bbk_sheet -mode mk > Makefile