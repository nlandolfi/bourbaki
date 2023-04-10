# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subrings.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

subrings.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > subrings.tex

subrings.pdf: ../../*.tex ../../trademark.pdf *.tex subrings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subrings.tex
	../../latexrun subrings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subrings.pdf
	open subrings.pdf

o:
	make open

terms: subrings.tex
	bbk_sheet -mode ts subrings.tex

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
	rm -f subrings.*
	make reset
	bbk_sheet -mode mk > Makefile