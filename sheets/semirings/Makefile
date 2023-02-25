# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make semirings.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

semirings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > semirings.tex

semirings.pdf: ../../*.tex ../../trademark.pdf *.tex semirings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode semirings.tex
	../../latexrun semirings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: semirings.pdf
	open semirings.pdf

o:
	make open

terms: semirings.tex
	bbk_sheet -mode ts semirings.tex

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
	rm -f semirings.*
	make reset
	bbk_sheet -mode mk > Makefile