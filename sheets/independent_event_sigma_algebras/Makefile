# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make independent_event_sigma_algebras.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

independent_event_sigma_algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > independent_event_sigma_algebras.tex

independent_event_sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex independent_event_sigma_algebras.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode independent_event_sigma_algebras.tex
	../../latexrun independent_event_sigma_algebras.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_event_sigma_algebras.pdf
	open independent_event_sigma_algebras.pdf

o:
	make open

terms: independent_event_sigma_algebras.tex
	bbk_sheet -mode ts independent_event_sigma_algebras.tex

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
	rm -f independent_event_sigma_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile